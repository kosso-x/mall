// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"mall/api/order/v1"
)

type IOrderV1 interface {
	OrderSubmit(ctx context.Context, req *v1.OrderSubmitReq) (res *v1.OrderSubmitRes, err error)
	OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
	OrderDelete(ctx context.Context, req *v1.OrderDeleteReq) (res *v1.OrderDeleteRes, err error)
	OrderCancel(ctx context.Context, req *v1.OrderCancelReq) (res *v1.OrderCancelRes, err error)
	OrderConfirm(ctx context.Context, req *v1.OrderConfirmReq) (res *v1.OrderConfirmRes, err error)
	OrderCount(ctx context.Context, req *v1.OrderCountReq) (res *v1.OrderCountRes, err error)
	OrderOrderCount(ctx context.Context, req *v1.OrderOrderCountReq) (res *v1.OrderOrderCountRes, err error)
	OrderExpress(ctx context.Context, req *v1.OrderExpressReq) (res *v1.OrderExpressRes, err error)
	OrderOrderGoods(ctx context.Context, req *v1.OrderOrderGoodsReq) (res *v1.OrderOrderGoodsRes, err error)
}
