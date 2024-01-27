package address

import (
	"context"

	v1 "mall/api/address/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AddressDelete(ctx context.Context, req *v1.AddressDeleteReq) (res *v1.AddressDeleteRes, err error) {
	service.Address().Delete(ctx, req.Id)

	res = &v1.AddressDeleteRes{
		Is_delete: 1,
	}

	return
}
