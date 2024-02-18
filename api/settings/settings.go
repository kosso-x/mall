// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package settings

import (
	"context"

	"mall/api/settings/v1"
)

type ISettingsV1 interface {
	ShowSettings(ctx context.Context, req *v1.ShowSettingsReq) (res *v1.ShowSettingsRes, err error)
	SaveSettings(ctx context.Context, req *v1.SaveSettingsReq) (res *v1.SaveSettingsRes, err error)
	SettingsDetail(ctx context.Context, req *v1.SettingsDetailReq) (res *v1.SettingsDetailRes, err error)
}
