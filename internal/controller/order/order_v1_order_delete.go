package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"
)

func (c *ControllerV1) OrderDelete(ctx context.Context, req *v1.OrderDeleteReq) (res *v1.OrderDeleteRes, err error) {
	var is_delete = 1
	err = service.Order().Delete(ctx, req.OrderId)
	if err != nil {
		is_delete = 0
		return
	}
	res = &v1.OrderDeleteRes{
		Is_delete: is_delete,
	}

	return
}
