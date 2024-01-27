package goods

import (
	"context"

	v1 "mall/api/goods/v1"
	"mall/internal/service"
)

func (c *ControllerV1) GoodsShare(ctx context.Context, req *v1.GoodsShareReq) (res *v1.GoodsShareRes, err error) {
	id := req.Id
	r, _ := service.Goods().Share(ctx, id)
	res = &v1.GoodsShareRes{
		GoodsShare: r,
	}

	return
}
