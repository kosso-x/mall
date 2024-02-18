package order

import (
	"context"
	"fmt"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/entity"
	"mall/internal/packed"
	"mall/internal/service"
	"math/rand"
	"time"

	"mall/internal/dao/tableOperation"

	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	OrderSelect      = []int{101, 102, 103, 201, 202, 203, 300, 301, 302, 303, 401}
	ToPaySelect      = []int{101, 801}
	ToDeliverySelect = []int{300}
	ToReceiveSelect  = []int{301}
	OrderTypeSelect  = 7
)

type (
	sOrder struct{}
)

func init() {
	service.RegisterOrder(New())
}

func New() service.IOrder {
	return &sOrder{}
}

func (s *sOrder) Submit(ctx context.Context, submit *model.OrderSubmit) (res *model.OrderInfo, err error) {
	/*
		1. 检测地址，没地址报错
		2. 检测订单，订单内没商品报错
		3. 检测库存，没有库存或者库存价格发生变化报错
		4. 生成商品清单
		5. 统计商品总价(纯商品价格)
		6. 订单价格计算(物流价格，使用红包，满减价格)
		7. 生成订单信息(order_sn生成)
		8. 开启事务，插入订单信息和订单商品
		9. 将商品信息录入数据库
	*/
	res = &model.OrderInfo{}
	if err = s.CheckAddress(ctx, submit.AddressId, res); err != nil {
		return nil, err
	}

	cart, err := s.CheckCart(ctx)
	if err != nil {
		return nil, err
	}

	if err = s.CheckGoods(ctx, cart, res); err != nil {
		return nil, err
	}

	res.Order_price += float32(submit.FreightPrice)
	res.Actual_price = res.Order_price - 0.00 // 减去红包金额，满减金额等
	res.Change_price = res.Order_price - 0.00
	res.Order_sn, err = s.GenerateOrderNumber(ctx)
	if err != nil {
		return nil, err
	}

	res.User_id = cmd.CurrentUser.Id
	res.Order_status = 101
	res.Freight_price = float32(submit.FreightPrice)
	res.Postscript = gbase64.EncodeString(submit.Postscript)
	res.Add_time = gconv.Int(time.Now().Unix())
	res.Offline_pay = submit.OfflinePay

	if tx, err := g.DB().Begin(ctx); err == nil {
		order_id, err := s.InertOrder(ctx, res)
		if err != nil {
			tx.Rollback()
		}
		res.Id = order_id

		if err = s.InsertOrderGoods(ctx, cart, order_id); err != nil {
			tx.Rollback()
		}

		if err = s.UpdateCart(ctx); err != nil {
			tx.Rollback()
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} else {
		return nil, err
	}

	return
}

func (s *sOrder) List(ctx context.Context, req *model.OrderListReq) (res model.OrderList, err error) {
	// var order []entity.HiolabsOrder
	res.TotalPages = packed.TotalPages(ctx, res.Count, req.Size)
	res.PageSize = req.Size
	res.CurrentPage = req.Page

	dao.
		HiolabsOrder.
		Ctx(ctx).
		Where("user_id = ? and is_delete = ? and order_type < ? and order_status in (?)", cmd.CurrentUser.Id, 0, OrderTypeSelect, OrderSelect).
		Limit((req.Page-1)*req.Size, req.Page*req.Size).
		OrderDesc("add_time").
		ScanAndCount(&res.Data, &res.Count, false)

	for _, data := range res.Data {
		dao.
			HiolabsOrderGoods.
			Ctx(ctx).
			Fields("id", "list_pic_url", "number").
			Where("user_id = ? and order_id = ? and is_delete = ?", cmd.CurrentUser.Id, data.Id, 0).
			Scan(&data.GoodsList)
		data.GoodsCount = len(data.GoodsList)
		value, found := tableOperation.OrderStatusEnum(data.OrderStatus)
		if !found {
			return res, gerror.New("get order status text error")
		}
		data.OrderStatusText = value
		data.HandleOption = *s.GetHandleOption(ctx, data.OrderStatus)
	}

	return
}

func (s *sOrder) Detail(ctx context.Context, order_id int) (res *model.OrderDetail, err error) {
	var (
		orderInfo    *model.OrderDetailInfo
		orderGoods   []model.HOrderGoods
		handleOption *model.HandleOption
		textCode     *model.TextCode
		goodsCount   = 0
	)
	// 订单信息
	dao.HiolabsOrder.Ctx(ctx).Where("user_id = ? and id = ?", cmd.CurrentUser.Id, order_id).Scan(&orderInfo)
	values := packed.GetFullRegion(ctx, &model.RegionIds{
		Country_id:  int(orderInfo.Country),
		Province_id: int(orderInfo.Province),
		City_id:     int(orderInfo.City),
		District_id: int(orderInfo.District),
		Full_region: "",
	})
	if err := gconv.Struct(values, orderInfo); err != nil {
		return nil, err
	}
	value, found := tableOperation.OrderStatusEnum(orderInfo.OrderStatus)
	if !found {
		return nil, gerror.New("get order status text error")
	}
	orderInfo.Order_status_text = value
	s.UpdateOrderStatus(ctx, orderInfo)

	// 物品信息
	dao.HiolabsOrderGoods.Ctx(ctx).Where("user_id = ? and order_id = ? and is_delete = ?", cmd.CurrentUser.Id, order_id, 0).Scan(&orderGoods)

	// 订单可操作按钮
	handleOption = s.GetHandleOption(ctx, int(orderInfo.OrderStatus))

	// textCode
	textCode = s.GetTextCode(ctx, int(orderInfo.OrderStatus))

	// 订单物品总数
	for _, g := range orderGoods {
		goodsCount += int(g.Number)
	}

	res = &model.OrderDetail{
		OrderInfo:    *orderInfo,
		OrderGoods:   orderGoods,
		HandleOption: *handleOption,
		TextCode:     *textCode,
		GoodsCount:   goodsCount,
	}

	return
}

func (s *sOrder) Delete(ctx context.Context, order_id int) (err error) {
	var order *entity.HiolabsOrder
	dao.HiolabsOrder.Ctx(ctx).Where("id = ?", order_id).Scan(&order)
	handleOption := s.GetHandleOption(ctx, int(order.OrderStatus))
	if !handleOption.Delete {
		return gerror.New("订单不能删除")
	} else {
		_, err = dao.HiolabsOrder.Ctx(ctx).Update(g.Map{
			"is_delete": 1,
		}, "id = ?", order_id)

		if err != nil {
			return
		}
	}

	return
}

func (s *sOrder) Cancel(ctx context.Context, order_id int) (err error) {
	var (
		order *entity.HiolabsOrder
		goods []entity.HiolabsOrderGoods
	)
	dao.HiolabsOrder.Ctx(ctx).Where("id = ?", order_id).Scan(&order)
	handleOption := s.GetHandleOption(ctx, int(order.OrderStatus))

	if !handleOption.Cancel {
		return gerror.New("订单不能取消")
	}

	if tx, err := g.DB().Begin(ctx); err == nil {
		//取消订单，还原库存
		dao.HiolabsOrderGoods.Ctx(ctx).Where("order_id = ? and user_id = ?", order_id, cmd.CurrentUser.Id).Scan(&goods)
		for _, gd := range goods {
			_, err = dao.HiolabsGoods.Ctx(ctx).Where("id = ?", gd.GoodsId).Increment("goods_number", gd.Number)
			if err != nil {
				tx.Rollback()
			}
			_, err = dao.HiolabsProduct.Ctx(ctx).Where("id = ?", gd.ProductId).Increment("goods_number", gd.Number)
			if err != nil {
				tx.Rollback()
			}
		}

		_, err = dao.HiolabsOrder.Ctx(ctx).Update(g.Map{
			"order_status": 102,
		}, "id = ?", order_id)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} else {
		return err
	}

	return
}

func (s *sOrder) Confirm(ctx context.Context, order_id int) (err error) {
	var order *entity.HiolabsOrder
	dao.HiolabsOrder.Ctx(ctx).Where("id = ?", order_id).Scan(&order)
	handleOption := s.GetHandleOption(ctx, int(order.OrderStatus))
	if !handleOption.Confirm {
		return gerror.New("订单不能确认")
	} else {
		_, err = dao.HiolabsOrder.Ctx(ctx).Update(g.Map{
			"order_status": 401,
			"confirm_time": gconv.Int(time.Now().Unix()),
		}, "id = ?", order_id)

		if err != nil {
			return
		}
	}

	return
}

func (s *sOrder) Count(ctx context.Context, showType int) (count int, err error) {
	status, err := s.GetOrderStatus(ctx, showType)
	if err != nil {
		return 0, err
	}

	count, err = dao.
		HiolabsOrder.
		Ctx(ctx).
		Where("user_id = ? and is_delete = ? and order_status in (?)", cmd.CurrentUser.Id, 0, status).
		Count()

	if err != nil {
		return 0, err
	}

	return
}

func (s *sOrder) OrderCount(ctx context.Context) (res *model.OrderCountInfo, err error) {
	r := dao.HiolabsOrder.Ctx(ctx).Where("user_id = ? and is_delete = ? and order_type < ? ", cmd.CurrentUser.Id, 0, OrderTypeSelect)

	toPayCount, err := r.Where("order_status in (?)", ToPaySelect).Count()
	if err != nil {
		return nil, err
	}
	toDeliveryCount, err := r.Where("order_status in (?)", ToDeliverySelect).Count()
	if err != nil {
		return nil, err
	}
	toReceiveCount, err := r.Where("order_status in (?)", ToReceiveSelect).Count()
	if err != nil {
		return nil, err
	}

	res = &model.OrderCountInfo{
		ToPay:      toPayCount,
		ToDelivery: toDeliveryCount,
		ToReceive:  toReceiveCount,
	}

	return
}

func (s *sOrder) Express(ctx context.Context) (err error) {
	return
}

func (s *sOrder) OrderGoods(ctx context.Context, order_id int) (res interface{}, err error) {
	if order_id > 0 {
		res, err = dao.
			HiolabsOrderGoods.
			Ctx(ctx).
			Where("user_id = ? and order_id = ? and is_delete = ?", cmd.CurrentUser.Id, order_id, 0).
			All()

		if err != nil {
			return
		}
	} else {
		res, err = dao.
			HiolabsCart.
			Ctx(ctx).
			Where("user_id = ? and checked = ? and is_delete = ? and is_fast = ?", cmd.CurrentUser.Id, 1, 0, 0).
			All()

		if err != nil {
			return
		}
	}

	return
}

func (s *sOrder) GetHandleOption(ctx context.Context, order_status int) (option *model.HandleOption) {
	var o model.HandleOption
	if order_status == 101 || order_status == 801 {
		o.Cancel = true
		o.Pay = true
	}
	// 如果订单被取消，则可删除
	if order_status == 102 || order_status == 103 {
		o.Delete = true
	}
	// 如果订单已付款，没有发货，则可退款操作
	// if (order_status == 201) {
	// 		handleOption.return = true;
	// }
	// 如果订单申请退款中，没有相关操作
	if order_status == 202 {
		o.CancelRefund = true
	}
	// 如果订单已经退款，则可删除
	if order_status == 300 {
		o.Delete = true
	}
	if order_status == 203 {
		o.Delete = true
	}
	// 如果订单已经发货，没有收货，则可收货操作,
	// 此时不能取消订单
	if order_status == 301 {
		o.Confirm = true
	}
	if order_status == 401 {
		o.Delete = true
	}

	option = &o

	return
}

func (s *sOrder) GetTextCode(ctx context.Context, order_status int) (text *model.TextCode) {
	var t model.TextCode
	if order_status == 101 {
		t.Pay = true
		t.Countdown = true
	}
	if order_status == 102 || order_status == 103 {
		t.Close = true
	}
	if order_status == 201 || order_status == 300 { //待发货
		t.Delivery = true
	}
	if order_status == 301 { //已发货
		t.Receive = true
	}
	if order_status == 401 {
		t.Success = true
	}
	text = &t
	return
}

func (s *sOrder) CheckAddress(ctx context.Context, address_id int, res *model.OrderInfo) (err error) {
	var address *entity.HiolabsAddress
	err = dao.HiolabsAddress.Ctx(ctx).Where("id = ?", address_id).Scan(&address)
	if err != nil {
		return
	}
	if address == nil {
		return gerror.New("请选择收货地址")
	}
	res.Consignee = address.Name
	res.Mobile = address.Mobile
	res.Province = address.ProvinceId
	res.City = address.CityId
	res.District = address.DistrictId
	res.Address = address.Address
	return
}

func (s *sOrder) CheckCart(ctx context.Context) (cart []entity.HiolabsCart, err error) {
	err = dao.HiolabsCart.Ctx(ctx).Where("user_id = ? and checked = ? and is_delete = ?", cmd.CurrentUser.Id, 1, 0).Scan(&cart)
	if err != nil {
		return
	}
	if cart == nil {
		return nil, gerror.New("请选择商品")
	}

	return
}

func (s *sOrder) CheckGoods(ctx context.Context, cart []entity.HiolabsCart, res *model.OrderInfo) (err error) {
	var product *entity.HiolabsProduct
	var printInfo string
	var totalPrice float32
	for i, c := range cart {
		dao.HiolabsProduct.Ctx(ctx).Where("id = ?", c.ProductId).Scan(&product)
		if c.Number > product.GoodsNumber {
			return gerror.New("库存不足，请重新下单")
		}
		if c.RetailPrice != c.AddPrice {
			return gerror.New("价格发生变化，请重新下单")
		}

		printInfo += gconv.String(i+1) + "、" + c.GoodsAka + "【" + gconv.String(c.Number) + "】"
		totalPrice += float32(c.Number) * float32(c.RetailPrice)
	}

	res.Print_info = printInfo
	res.Goods_price = totalPrice

	return
}

func (s *sOrder) GenerateOrderNumber(ctx context.Context) (sn string, err error) {
	now := time.Now()
	first_half := fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	last_half := gconv.String(rand.Intn(999999) + 100000)
	sn = first_half + last_half
	return
}

func (s *sOrder) InertOrder(ctx context.Context, order *model.OrderInfo) (order_id int, err error) {
	id, err := dao.HiolabsOrder.Ctx(ctx).InsertAndGetId(gconv.Map(order))
	if err != nil {
		return 0, err
	}
	order_id = gconv.Int(id)

	return
}

func (s *sOrder) InsertOrderGoods(ctx context.Context, cart []entity.HiolabsCart, order_id int) (err error) {
	for _, c := range cart {
		dao.HiolabsOrderGoods.Ctx(ctx).Insert(g.Map{
			"user_id":                      cmd.CurrentUser.Id,
			"order_id":                     order_id,
			"goods_id":                     c.GoodsId,
			"product_id":                   c.ProductId,
			"goods_name":                   c.GoodsName,
			"goods_aka":                    c.GoodsAka,
			"list_pic_url":                 c.ListPicUrl,
			"retail_price":                 c.RetailPrice,
			"number":                       c.Number,
			"goods_specifition_name_value": c.GoodsSpecifitionNameValue,
			"goods_specifition_ids":        c.GoodsSpecifitionIds,
		})
	}
	return
}

func (s *sOrder) UpdateCart(ctx context.Context) (err error) {
	_, err = dao.HiolabsCart.Ctx(ctx).Update("is_delete='1'", "user_id = ? and checked = ? and is_delete = ?", cmd.CurrentUser.Id, 1, 0)
	return
}

func (s *sOrder) UpdateOrderStatus(ctx context.Context, orderInfo *model.OrderDetailInfo) (err error) {
	if orderInfo.OrderStatus == 101 || orderInfo.OrderStatus == 801 {
		orderInfo.Final_pay_time = gconv.Int(orderInfo.AddTime) + 24*60*60 //支付倒计时2小时
		if orderInfo.Final_pay_time < gconv.Int(time.Now().Unix()) {
			//超过时间不支付，更新订单状态为取消
			dao.HiolabsOrder.Ctx(ctx).Update(g.Map{
				"order_status": 102,
			}, "id = ?", orderInfo.Id)

			orderInfo.OrderStatus = 102
		}
	}

	return
}

func (s *sOrder) GetOrderStatus(ctx context.Context, showType int) (status []int, err error) {
	switch showType {
	case 0:
		status = []int{101, 102, 103, 201, 202, 203, 300, 301, 302, 303, 401}
	case 1:
		status = []int{101}
	case 2:
		status = []int{300}
	case 3:
		status = []int{301}
	case 4:
		status = []int{302, 303}
	}
	return
}
