package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) CartAdd(ctx context.Context, req *v1.CartAddReq) (res *v1.CartAddRes, err error) {
	req_body := &model.CartAddReq{
		AddType:   req.AddType,
		GoodsId:   req.GoodsId,
		Number:    req.Number,
		ProductId: req.ProductId,
	}

	r, _ := service.Cart().Add(ctx, req_body)

	res = &v1.CartAddRes{
		CartAddRes: &r,
	}

	return
}
