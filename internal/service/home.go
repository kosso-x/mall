// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mall/internal/model"
)

type (
	IHome interface {
		Info(ctx context.Context) (res *model.Info, err error)
		InfoShortCategory(ctx context.Context) (res []model.ShortCategory, err error)
		InfoShortAd(ctx context.Context) (res []model.ShortAd, err error)
		InfoShortNotice(ctx context.Context) (res []model.ShortNotice, err error)
		InfoCategoryGood(ctx context.Context) (res []model.CategoryGood, err error)
		InfoGoods(ctx context.Context, category_id int) (res []model.GoodsList, err error)
		InfoCartCount(ctx context.Context) (res int, err error)
		Encode(ctx context.Context, GoodsId int) (text string, err error)
	}
)

var (
	localHome IHome
)

func Home() IHome {
	if localHome == nil {
		panic("implement not found for interface IHome, forgot register?")
	}
	return localHome
}

func RegisterHome(i IHome) {
	localHome = i
}
