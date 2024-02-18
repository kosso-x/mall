package home

import (
	"context"

	v1 "mall/api/home/v1"
	"mall/internal/service"
)

func (c *ControllerV1) GetBase64(ctx context.Context, req *v1.GetBase64Req) (res *v1.GetBase64Res, err error) {
	r, err := service.Home().Encode(ctx, req.GoodsId)
	if err != nil {
		return
	}

	res = &v1.GetBase64Res{
		Ciphertext: r,
	}
	return
}
