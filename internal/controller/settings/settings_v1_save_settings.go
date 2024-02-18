package settings

import (
	"context"

	v1 "mall/api/settings/v1"
	"mall/internal/service"
)

func (c *ControllerV1) SaveSettings(ctx context.Context, req *v1.SaveSettingsReq) (res *v1.SaveSettingsRes, err error) {
	var saved = 1
	err = service.Settings().Save(ctx, req.SaveSettings)
	if err != nil {
		saved = 0
		return
	}

	res = &v1.SaveSettingsRes{
		Saved: saved,
	}

	return
}
