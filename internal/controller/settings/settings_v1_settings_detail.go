package settings

import (
	"context"

	v1 "mall/api/settings/v1"
	"mall/internal/service"
)

func (c *ControllerV1) SettingsDetail(ctx context.Context, req *v1.SettingsDetailReq) (res *v1.SettingsDetailRes, err error) {
	r, err := service.Settings().UserDetail(ctx)
	if err != nil {
		return
	}

	res = &v1.SettingsDetailRes{
		UserDetail: r,
	}
	return
}
