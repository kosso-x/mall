package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (res *v1.CartUpdateRes, err error) {
	update := &model.UpdateCart{
		Id:        req.Id,
		Number:    req.Number,
		ProductId: req.ProductId,
	}

	r, err := service.Cart().Update(ctx, update)
	res = &v1.CartUpdateRes{
		CartAddRes: &r,
	}

	return
}
