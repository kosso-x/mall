package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderCancel(ctx context.Context, req *v1.OrderCancelReq) (res *v1.OrderCancelRes, err error) {
	var canceled = 1
	err = service.Order().Cancel(ctx, req.OrderId)
	if err != nil {
		canceled = 0
	}
	res = &v1.OrderCancelRes{
		Canceled: canceled,
	}

	return
}
