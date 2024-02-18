package order

import (
	"context"

	v1 "mall/api/order/v1"
	"mall/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) OrderOrderGoods(ctx context.Context, req *v1.OrderOrderGoodsReq) (res *v1.OrderOrderGoodsRes, err error) {
	r, err := service.Order().OrderGoods(ctx, req.OrderId)
	if err != nil {
		return
	}

	g.Dump(r)
	g.Dump("RRRRRRRRRRRRRRRRRRRRRR")
	res = &v1.OrderOrderGoodsRes{
		Res: r,
	}
	return
}
