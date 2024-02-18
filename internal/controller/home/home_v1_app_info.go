package home

import (
	"context"

	v1 "mall/api/home/v1"
	"mall/internal/service"
)

func (c *ControllerV1) AppInfo(ctx context.Context, req *v1.AppInfoReq) (res *v1.AppInfoRes, err error) {
	r, err := service.Home().Info(ctx)

	res = &v1.AppInfoRes{
		Channel:      &r.Channel,
		Banner:       &r.Banner,
		Notice:       &r.Notice,
		CategoryList: &r.CategoryList,
		CartCount:    &r.CartCount,
	}

	return
}
