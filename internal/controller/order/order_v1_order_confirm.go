package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderConfirm(ctx context.Context, req *v1.OrderConfirmReq) (res *v1.OrderConfirmRes, err error) {
	var confirmed = 1
	err = service.Order().Confirm(ctx, req.OrderId)
	if err != nil {
		confirmed = 0
	}
	res = &v1.OrderConfirmRes{
		Confirmed: confirmed,
	}

	return
}
