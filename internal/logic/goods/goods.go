package goods

import (
	"context"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/entity"
	"mall/internal/service"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sGoods struct{}
)

func init() {
	service.RegisterGoods(New())
}

func New() service.IGoods {
	return &sGoods{}
}

func (s *sGoods) Count(ctx context.Context) (count int, err error) {
	count, err = dao.HiolabsGoods.Ctx(ctx).Where("is_delete = ? and is_on_sale = ?", 0, 1).Count()
	if err != nil {
		return 0, gerror.New("get goods count error")
	}

	return
}

func (s *sGoods) Detail(ctx context.Context, id int) (detail *model.GoodsDetail, err error) {
	goods, _ := s.GetGoods(ctx, id)
	gallerys, _ := s.GetGallery(ctx, id)
	gs, _ := s.GetSpecification(ctx, id)
	products, _ := s.GetProduct(ctx, id)

	detail = &model.GoodsDetail{
		Info:              *goods,
		Gallery:           gallerys,
		SpecificationList: *gs,
		ProductList:       products,
	}

	return
}

func (s *sGoods) List(ctx context.Context) (res model.GoodsListInfo, err error) {
	err = dao.HiolabsGoods.Ctx(ctx).Where("is_on_sale = ? and is_delete = ?", 1, 0).OrderAsc("sort_order").Scan(&res)
	if err != nil {
		return nil, gerror.New("get goods error")
	}

	return
}

func (s *sGoods) Share(ctx context.Context, id int) (share *model.GoodsShare, err error) {
	err = dao.HiolabsGoods.Ctx(ctx).Fields("name", "retail_price").Where("id = ?", id).Scan(&share)
	if err != nil {
		return nil, gerror.New("get goods share err")
	}

	return
}

func (s *sGoods) GetGoods(ctx context.Context, id int) (goods *entity.HiolabsGoods, err error) {
	err = dao.HiolabsGoods.Ctx(ctx).Where("id = ? and is_delete = ?", id, 0).Scan(&goods)
	if err != nil {
		return nil, gerror.New("get goods error")
	}

	return
}

func (s *sGoods) GetGallery(ctx context.Context, id int) (gallerys []entity.HiolabsGoodsGallery, err error) {
	err = dao.HiolabsGoodsGallery.Ctx(ctx).Where("goods_id = ? and is_delete = ?", id, 0).OrderAsc("sort_order").Scan(&gallerys)
	if err != nil {
		return nil, gerror.New("get goods gallery error")
	}

	return
}

func (s *sGoods) GetSpecification(ctx context.Context, id int) (gs *model.GoodsSpecification, err error) {
	var (
		goods_spec model.ValueList
		spec       entity.HiolabsSpecification
		goods      entity.HiolabsGoods
	)
	dao.HiolabsGoodsSpecification.Ctx(ctx).Where("goods_id = ? and is_delete = ?", id, 0).Scan(&goods_spec)
	dao.HiolabsSpecification.Ctx(ctx).Where("id = ?", goods_spec.Specification_id).Scan(&spec)
	dao.HiolabsGoods.Ctx(ctx).Where("id = ? and is_delete = ?", id, 0).Scan(&goods)
	goods_spec.Goods_number = int(goods.GoodsNumber)

	var values []model.ValueList
	if err := gconv.Scan(goods_spec, &values); err != nil {
		return nil, gerror.New("get specification error")
	}

	gs = &model.GoodsSpecification{
		Specification_id: int(spec.Id),
		Name:             spec.Name,
		ValueList:        values,
	}

	return
}

func (s *sGoods) GetProduct(ctx context.Context, id int) (products []entity.HiolabsProduct, err error) {
	dao.HiolabsProduct.Ctx(ctx).Where("goods_id = ? AND is_delete = ?", id, 0).Scan(&products)
	if err != nil {
		return nil, gerror.New("get goods products error")
	}

	return
}
