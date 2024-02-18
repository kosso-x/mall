package service

import (
	"context"
	"mall/internal/model"
)

type (
	IGoods interface {
		Count(ctx context.Context) (count int, err error)
		Detail(ctx context.Context, id int) (detail *model.GoodsDetail, err error)
		List(ctx context.Context) (res model.GoodsListInfo, err error)
		Share(ctx context.Context, id int) (share *model.GoodsShare, err error)
		GetGoods(ctx context.Context, id int) (goods *model.HGoods, err error)
		GetGallery(ctx context.Context, id int) (gallerys []model.HGoodsGallery, err error)
		GetSpecification(ctx context.Context, id int) (values *model.GoodsSpecification, err error)
		GetProduct(ctx context.Context, id int) (products []model.HProduct, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
