package goods

import (
	"context"

	v1 "mall/api/goods/v1"
	"mall/internal/service"
)

func (c *ControllerV1) GoodsList(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error) {
	r, err := service.Goods().List(ctx)
	res = &v1.GoodsListRes{
		GoodsListInfo: &r,
	}

	return
}
