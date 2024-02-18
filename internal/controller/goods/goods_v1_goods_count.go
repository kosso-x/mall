package goods

import (
	"context"

	v1 "mall/api/goods/v1"
	"mall/internal/service"
)

func (c *ControllerV1) GoodsCount(ctx context.Context, req *v1.GoodsCountReq) (res *v1.GoodsCountRes, err error) {
	count, err := service.Goods().Count(ctx)
	res = &v1.GoodsCountRes{
		GoodsCount: count,
	}

	return
}
