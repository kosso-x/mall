package address

import (
	"context"

	v1 "mall/api/address/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AddressSave(ctx context.Context, req *v1.AddressSaveReq) (res *v1.AddressSaveRes, err error) {
	address := req.AddressSave
	r, _ := service.Address().Save(ctx, address)
	res = &v1.AddressSaveRes{
		AddressSaveRes: r,
	}
	return
}
