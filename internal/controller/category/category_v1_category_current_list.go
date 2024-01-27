package category

import (
	"context"

	v1 "mall/api/category/v1"
	"mall/internal/service"
)

func (c *ControllerV1) CategoryCurrentList(ctx context.Context, req *v1.CategoryCurrentListReq) (res *v1.CategoryCurrentListRes, err error) {
	r, err := service.Category().CurrentList(ctx, req.Id, req.Page, req.Size)
	res = &v1.CategoryCurrentListRes{
		CategoryCurrentList: &r,
	}

	if err != nil {
		return nil, err
	}

	return
}
