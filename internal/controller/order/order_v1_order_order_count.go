package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderOrderCount(ctx context.Context, req *v1.OrderOrderCountReq) (res *v1.OrderOrderCountRes, err error) {
	r, err := service.Order().OrderCount(ctx)
	if err != nil {
		return
	}
	res = &v1.OrderOrderCountRes{
		OrderCountInfo: r,
	}

	return
}
