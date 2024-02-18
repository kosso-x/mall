package footprint

import (
	"context"

	v1 "mall/api/footprint/v1"
	"mall/internal/service"
)

func (c *ControllerV1) FootprintList(ctx context.Context, req *v1.FootprintListReq) (res *v1.FootprintListRes, err error) {
	r, err := service.Footprint().List(ctx, req.Page, req.Size)
	if err != nil {
		return
	}
	res = &v1.FootprintListRes{
		FootprintList: r,
	}

	return
}
