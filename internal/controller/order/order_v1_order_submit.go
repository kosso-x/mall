package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) OrderSubmit(ctx context.Context, req *v1.OrderSubmitReq) (res *v1.OrderSubmitRes, err error) {
	submit := &model.OrderSubmit{
		AddressId:    req.AddressId,
		Postscript:   req.Postscript,
		FreightPrice: req.FreightPrice,
		ActualPrice:  req.ActualPrice,
		OfflinePay:   req.OfflinePay,
	}
	r, err := service.Order().Submit(ctx, submit)
	res = &v1.OrderSubmitRes{
		OrderInfo: r,
	}

	return
}
