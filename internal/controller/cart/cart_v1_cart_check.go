package cart

import (
	"context"

	v1 "mall/api/cart/v1"
	"mall/internal/model"
	"mall/internal/service"
)

func (c *ControllerV1) CartCheck(ctx context.Context, req *v1.CartCheckReq) (res *v1.CartCheckRes, err error) {
	check := &model.CheckedCart{
		ProductIds: req.ProductIds,
		IsChecked:  req.IsChecked,
	}
	r, err := service.Cart().Check(ctx, check)

	res = &v1.CartCheckRes{
		CartAddRes: &r,
	}

	return
}
