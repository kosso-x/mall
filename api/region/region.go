// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package region

import (
	"context"

	"mall/api/region/v1"
)

type IRegionV1 interface {
	RegionList(ctx context.Context, req *v1.RegionListReq) (res *v1.RegionListRes, err error)
}
