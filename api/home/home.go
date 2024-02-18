// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package home

import (
	"context"

	"mall/api/home/v1"
)

type IHomeV1 interface {
	AppInfo(ctx context.Context, req *v1.AppInfoReq) (res *v1.AppInfoRes, err error)
	GetBase64(ctx context.Context, req *v1.GetBase64Req) (res *v1.GetBase64Res, err error)
}
