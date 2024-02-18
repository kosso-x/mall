package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	r, err := service.Order().Detail(ctx, req.OrderId)
	res = &v1.OrderDetailRes{
		OrderDetail: r,
	}

	return
}
