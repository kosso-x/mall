package address

import (
	"context"

	v1 "mall/api/address/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AddressGet(ctx context.Context, req *v1.AddressGetReq) (res *v1.AddressGetRes, err error) {
	r, err := service.Address().List(ctx)
	res = &v1.AddressGetRes{
		AddressList: &r,
	}

	return
}
