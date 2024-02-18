package settings

import (
	"context"

	v1 "mall/api/settings/v1"
	"mall/internal/service"
)

func (c *ControllerV1) ShowSettings(ctx context.Context, req *v1.ShowSettingsReq) (res *v1.ShowSettingsRes, err error) {
	r, err := service.Settings().Show(ctx)
	if err != nil {
		return
	}

	res = &v1.ShowSettingsRes{
		ShowSettings: r,
	}
	return
}
