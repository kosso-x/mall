package address

import (
	"context"

	v1 "mall/api/address/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AddressDetail(ctx context.Context, req *v1.AddressDetailReq) (res *v1.AddressDetailRes, err error) {
	r, err := service.Address().Detail(ctx, req.Id)
	res = &v1.AddressDetailRes{
		AddressDetail: r,
	}

	return
}
