package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	list_req := &model.OrderListReq{
		ShowType: req.ShowType,
		Size:     req.Size,
		Page:     req.Page,
	}

	r, err := service.Order().List(ctx, list_req)
	res = &v1.OrderListRes{
		OrderList: &r,
	}

	return
}
