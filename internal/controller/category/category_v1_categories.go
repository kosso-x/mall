package category

import (
	"context"

	v1 "mall/api/category/v1"
	"mall/internal/service"
)

func (c *ControllerV1) Categories(ctx context.Context, req *v1.CategoriesReq) (res *v1.CategoriesRes, err error) {
	r, err := service.Category().List(ctx)
	res = &v1.CategoriesRes{
		CategoryList: &r,
	}

	if err != nil {
		return nil, err
	}

	return
}
