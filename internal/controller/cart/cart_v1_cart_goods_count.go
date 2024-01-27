package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/service"
)

func (c *ControllerV1) CartGoodsCount(ctx context.Context, req *v1.CartGoodsCountReq) (res *v1.CartGoodsCountRes, err error) {
	r, err := service.Cart().GoodsCount(ctx)
	res = &v1.CartGoodsCountRes{
		CartGoodsCount: &r,
	}

	return
}
