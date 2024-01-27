package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/service"
)

func (c *ControllerV1) CartIndex(ctx context.Context, req *v1.CartIndexReq) (res *v1.CartIndexRes, err error) {
	r, _ := service.Cart().List(ctx)
	res = &v1.CartIndexRes{
		CartAddRes: &r,
	}

	return
}
