package cart

import (
	"context"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/entity"
	"mall/internal/service"
	"time"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var user = cmd.CurrentUser

type (
	sCart struct{}
)

func init() {
	service.RegisterCart(New())
}

func New() service.ICart {
	return &sCart{}
}

func (s *sCart) Add(ctx context.Context, req *model.CartAddReq) (res model.CartAddRes, err error) {
	/*
		查看有没有库存
		查看有没有产品
		查看当前购物车里的该产品
		修改购物车产品
	*/

	count_check, goods, err := s.CheckCount(ctx, req.GoodsId)
	if !count_check {
		return
	}

	product_check, product, err := s.CheckProduct(ctx, req.ProductId)
	if !product_check {
		return
	}

	s.CurrentCart(ctx, product, goods)
	res, err = s.List(ctx)

	return
}

func (s *sCart) List(ctx context.Context) (res model.CartAddRes, err error) {
	var carts []model.CartList
	dao.HiolabsCart.Ctx(ctx).Where("user_id = ? and is_delete = ? and is_fast = ?", cmd.CurrentUser.Id, 0, 0).Scan(&carts)
	for i, _ := range carts {
		carts[i].Weight_count = carts[i].Number * carts[i].Goods_weight
	}
	res.CartList = carts

	for _, c := range carts {
		res.CartTotal.GoodsCount = res.CartTotal.GoodsCount + c.Number
		res.CartTotal.GoodsAmount = res.CartTotal.GoodsAmount + float32(c.Number)*c.Retail_price
		if c.Checked == 1 {
			res.CartTotal.CheckedGoodsCount = res.CartTotal.CheckedGoodsCount + c.Number
			res.CartTotal.CheckedGoodsAmount = res.CartTotal.CheckedGoodsAmount + float32(c.Number)*c.Retail_price
		}

		res.CartTotal.User_id = cmd.CurrentUser.Id
		res.CartTotal.NumberChange = 0
	}

	if err != nil {
		return res, gerror.New("get current cart error")
	}

	return
}

func (s *sCart) Update(ctx context.Context, update *model.UpdateCart) (res model.CartAddRes, err error) {
	/*
		找到该用户的购物车内的该商品信息
		更新
	*/
	dao.HiolabsCart.Ctx(ctx).Update(g.Map{
		"number": update.Number,
	}, "id = ? and product_id = ? and user_id = ?", update.Id, update.ProductId, cmd.CurrentUser.Id)

	res, err = s.List(ctx)
	return
}

func (s *sCart) Delete(ctx context.Context, ProductIds []int) (res model.CartAddRes, err error) {
	_, err = dao.HiolabsCart.Ctx(ctx).Delete("user_id = ? and product_id in (?)", cmd.CurrentUser.Id, ProductIds)
	if err != nil {
		return res, err
	}

	res, err = s.List(ctx)

	return
}

func (s *sCart) Check(ctx context.Context, check *model.CheckedCart) (res model.CartAddRes, err error) {
	/*
		找到该用户的购物车内的该商品信息
		更新checked
	*/
	dao.HiolabsCart.Ctx(ctx).Update(g.Map{
		"checked": check.IsChecked,
	}, "product_id in (?) and user_id = ?", check.ProductIds, cmd.CurrentUser.Id)

	res, err = s.List(ctx)
	return
}

func (s *sCart) GoodsCount(ctx context.Context) (res model.CartGoodsCount, err error) {
	var carts []model.CartList
	dao.HiolabsCart.Ctx(ctx).Where("user_id = ? and is_delete = ? and is_fast = ?", cmd.CurrentUser.Id, 0, 0).Scan(&carts)

	for _, c := range carts {
		res.CartTotal.GoodsCount = res.CartTotal.GoodsCount + c.Number
	}

	return
}

func (s *sCart) CheckOut(ctx context.Context, checkout *model.CartCheckoutReq) (res model.CartCheckout, err error) {
	carts, err := s.List(ctx)
	var (
		freight      float32
		freight_temp *entity.HiolabsFreightTemplate
	)
	for _, v := range carts.CartList {
		if v.Checked == 1 {
			dao.HiolabsFreightTemplate.Ctx(ctx).Where("id = ?", v.Freight_template_id).Scan(&freight_temp)
			freight = freight + (float32(freight_temp.PackagePrice) * float32(v.Number))
		}
	}

	res = model.CartCheckout{
		CheckedAddress:   0,
		FreightPrice:     freight,
		CheckedGoodsList: carts.CartList,
		GoodsTotalPrice:  carts.CartTotal.GoodsAmount,
		OrderTotalPrice:  carts.CartTotal.CheckedGoodsAmount,
		ActualPrice:      carts.CartTotal.CheckedGoodsAmount,
		GoodsCount:       carts.CartTotal.GoodsCount,
		OutStock:         0,
		NumberChange:     0,
	}

	return
}

func (s *sCart) CheckCount(ctx context.Context, GoodsId string) (checked bool, goods *entity.HiolabsGoods, err error) {
	err = dao.HiolabsGoods.Ctx(ctx).Where("id = ?", GoodsId).Scan(&goods)

	if err != nil {
		checked = false
		return
	}

	if (goods.IsOnSale == 0) || goods.GoodsNumber > 0 {
		checked = true
		err = nil
	} else {
		checked = false
		err = gerror.New("goods out of stock")
	}

	return
}

func (s *sCart) CheckProduct(ctx context.Context, ProductId int) (checked bool, product *entity.HiolabsProduct, err error) {
	err = dao.HiolabsProduct.Ctx(ctx).Where("id = ?", ProductId).Scan(&product)

	if err != nil {
		checked = false

		return
	}

	if (product.IsOnSale == 0) || product.GoodsNumber > 0 {
		checked = true
		err = nil
	} else {
		checked = false
		err = gerror.New("product out of stock")
	}

	return
}

func (s *sCart) CurrentCart(ctx context.Context, product *entity.HiolabsProduct, goods *entity.HiolabsGoods) (cart entity.HiolabsCart, err error) {
	var spec []entity.HiolabsGoodsSpecification
	dao.HiolabsCart.Ctx(ctx).Where("user_id = ? and product_id = ? and is_delete = ? and is_fast = ?", cmd.CurrentUser.Id, product.Id, 0, 0).Scan(&cart)
	// 用户购物车中存在该商品就加一，不存在就添加

	if cart.Id != 0 {
		dao.HiolabsCart.Ctx(ctx).Update(g.Map{
			"number": cart.Number + 1,
		}, "id", cart.Id)
		dao.HiolabsCart.Ctx(ctx).Where("id = ?", cart.Id).Scan(&cart)
	} else {
		var value string
		for _, s := range spec {
			value = value + s.Value
		}
		dao.HiolabsSpecification.Ctx(ctx).Where("goods_id = ?", goods.Id).Scan(&spec)
		dao.HiolabsCart.Ctx(ctx).Insert(g.Map{
			"user_id":                      cmd.CurrentUser.Id,
			"goods_id":                     product.GoodsId,
			"goods_sn":                     product.GoodsSn,
			"product_id":                   product.Id,
			"goods_name":                   product.GoodsName,
			"goods_aka":                    product.GoodsName,
			"goods_weight":                 product.GoodsWeight,
			"add_price":                    product.RetailPrice,
			"retail_price":                 product.RetailPrice,
			"number":                       1,
			"goods_specifition_name_value": value,
			"goods_specifition_ids":        product.GoodsSpecificationIds,
			"checked":                      1,
			"list_pic_url":                 goods.ListPicUrl,
			"freight_template_id":          goods.FreightTemplateId,
			"is_on_sale":                   product.IsOnSale,
			"add_time":                     time.Now().Unix(),
			"is_fast":                      0,
			"is_delete":                    product.IsDelete,
		})
		dao.HiolabsCart.Ctx(ctx).Where("user_id = ? and product_id = ? and is_delete = ? and is_fast = ?", cmd.CurrentUser.Id, product.Id, 0, 0).Scan(&cart)
	}
	return
}
