package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderCount(ctx context.Context, req *v1.OrderCountReq) (res *v1.OrderCountRes, err error) {
	count, err := service.Order().Count(ctx, req.ShowType)
	if err != nil {
		return
	}

	res = &v1.OrderCountRes{
		AllCount: count,
	}

	return
}
