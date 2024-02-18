package footprint

import (
	"context"

	v1 "mall/api/footprint/v1"
	"mall/internal/service"
)

func (c *ControllerV1) FootprintDelete(ctx context.Context, req *v1.FootprintDeleteReq) (res *v1.FootprintDeleteRes, err error) {
	var is_delete = 1
	err = service.Footprint().Delete(ctx, req.FootprintId)
	if err != nil {
		is_delete = 0
		return
	}

	res = &v1.FootprintDeleteRes{
		IsDelete: is_delete,
	}

	return
}
