// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mall/internal/model"
	"mall/internal/model/entity"
)

type (
	ICart interface {
		Add(ctx context.Context, req *model.CartAddReq) (res model.CartAddRes, err error)
		List(ctx context.Context) (res model.CartAddRes, err error)
		Update(ctx context.Context, update *model.UpdateCart) (res model.CartAddRes, err error)
		Delete(ctx context.Context, ProductIds []int) (res model.CartAddRes, err error)
		Check(ctx context.Context, check *model.CheckedCart) (res model.CartAddRes, err error)
		GoodsCount(ctx context.Context) (res model.CartGoodsCount, err error)
		CheckOut(ctx context.Context, checkout *model.CartCheckoutReq) (res model.CartCheckout, err error)

		// 数据库查询方法
		// 查询商品库存数量
		CheckCount(ctx context.Context, GoodsId string) (checked bool, goods *entity.HiolabsGoods, err error)
		// 查询商品规格是否可以购买
		CheckProduct(ctx context.Context, ProductId int) (checked bool, product *entity.HiolabsProduct, err error)
		// 当前用户购物车状态
		CurrentCart(ctx context.Context, product *entity.HiolabsProduct, goods *entity.HiolabsGoods) (cart entity.HiolabsCart, err error)
	}
)

var (
	localCart ICart
)

func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

func RegisterCart(i ICart) {
	localCart = i
}
