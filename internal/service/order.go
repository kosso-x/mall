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
	IOrder interface{
		Submit(ctx context.Context, submit *model.OrderSubmit) (res *model.OrderInfo, err error)
		List(ctx context.Context, req *model.OrderListReq) (res model.OrderList, err error)
		Detail(ctx context.Context, order_id int) (res *model.OrderDetail, err error)
		Delete(ctx context.Context, order_id int) (err error)
		Cancel(ctx context.Context, order_id int) (err error)
		Confirm(ctx context.Context, order_id int) (err error)
		Count(ctx context.Context, showType int) (count int, err error)
		OrderCount(ctx context.Context) (res *model.OrderCountInfo, err error)
		Express(ctx context.Context) (err error)
		OrderGoods(ctx context.Context, order_id int) (res interface{}, err error)
		GetHandleOption(ctx context.Context, order_status int) (option *model.HandleOption)
		GetTextCode(ctx context.Context, order_status int) (text *model.TextCode)
		CheckAddress(ctx context.Context, address_id int, res *model.OrderInfo) (err error)
		CheckCart(ctx context.Context) (cart []entity.HiolabsCart, err error)
		CheckGoods(ctx context.Context, cart []entity.HiolabsCart, res *model.OrderInfo) (err error)
		GenerateOrderNumber(ctx context.Context) (sn string, err error)
		InertOrder(ctx context.Context, order *model.OrderInfo) (order_id int, err error)
		InsertOrderGoods(ctx context.Context, cart []entity.HiolabsCart, order_id int) (err error)
		UpdateCart(ctx context.Context) (err error)
		UpdateOrderStatus(ctx context.Context, orderInfo *model.OrderDetailInfo) (err error)
		GetOrderStatus(ctx context.Context, showType int) (status []int, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
