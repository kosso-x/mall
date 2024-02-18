package order

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"mall/api/order/v1"
)

func (c *ControllerV1) OrderExpress(ctx context.Context, req *v1.OrderExpressReq) (res *v1.OrderExpressRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
