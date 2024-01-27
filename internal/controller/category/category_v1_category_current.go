package category

import (
	"context"

	v1 "mall/api/category/v1"
	"mall/internal/service"
)

func (c *ControllerV1) CategoryCurrent(ctx context.Context, req *v1.CategoryCurrentReq) (res *v1.CategoryCurrentRes, err error) {
	r, err := service.Category().One(ctx, req.Id)
	res = &v1.CategoryCurrentRes{
		CurrentCategory: &r,
	}

	if err != nil {
		return nil, err
	}
	return
}
