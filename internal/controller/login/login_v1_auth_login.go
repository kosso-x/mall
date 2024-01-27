package login

import (
	"context"

	v1 "mall/api/login/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	r, err := service.Login().AuthLogin(ctx, req.Code)

	res = &v1.AuthLoginRes{
		AuthLogin: r,
	}

	return
}
