package region

import (
	"context"

	v1 "mall/api/region/v1"
	"mall/internal/service"
)

func (c *ControllerV1) RegionList(ctx context.Context, req *v1.RegionListReq) (res *v1.RegionListRes, err error) {
	r, err := service.Region().List(ctx, req.ParentId)
	res = &v1.RegionListRes{
		RegionList: r,
	}

	return
}
