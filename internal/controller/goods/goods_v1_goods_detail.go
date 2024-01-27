package goods

import (
	"context"

	v1 "mall/api/goods/v1"
	"mall/internal/service"
)

func (c *ControllerV1) GoodsDetail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error) {
	id := req.Id
	r, err := service.Goods().Detail(ctx, id)
	res = &v1.GoodsDetailRes{
		GoodsDetail: r,
	}

	return
}
