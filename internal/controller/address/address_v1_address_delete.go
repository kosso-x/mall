package address

import (
	"context"

	v1 "mall/api/address/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AddressDelete(ctx context.Context, req *v1.AddressDeleteReq) (res *v1.AddressDeleteRes, err error) {
	var is_delete = 1
	err = service.Address().Delete(ctx, req.Id)
	if err != nil {
		is_delete = 0
		return
	}

	res = &v1.AddressDeleteRes{
		Is_delete: is_delete,
	}

	return
}
