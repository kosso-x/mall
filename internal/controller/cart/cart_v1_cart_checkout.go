package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (res *v1.CartCheckoutRes, err error) {
	checkout := &model.CartCheckoutReq{
		AddressId: req.AddressId,
		AddType:   req.AddType,
		OrderFrom: req.OrderFrom,
		Type:      req.Type,
	}

	r, err := service.Cart().CheckOut(ctx, checkout)
	res = &v1.CartCheckoutRes{
		CartCheckout: &r,
	}

	return
}
