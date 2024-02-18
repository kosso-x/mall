// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package footprint

import (
	"context"

	"mall/api/footprint/v1"
)

type IFootprintV1 interface {
	FootprintList(ctx context.Context, req *v1.FootprintListReq) (res *v1.FootprintListRes, err error)
	FootprintDelete(ctx context.Context, req *v1.FootprintDeleteReq) (res *v1.FootprintDeleteRes, err error)
}
