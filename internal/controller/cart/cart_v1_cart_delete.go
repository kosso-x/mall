package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/service"
)

func (c *ControllerV1) CartDelete(ctx context.Context, req *v1.CartDeleteReq) (res *v1.CartDeleteRes, err error) {
	product_id := req.ProductIds
	r, err := service.Cart().Delete(ctx, product_id)
	service.Cart().Delete(ctx, product_id)
	res = &v1.CartDeleteRes{
		CartAddRes: &r,
	}

	return
}
