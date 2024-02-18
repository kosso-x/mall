package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ShowSettingsReq struct {
	g.Meta `path:"/showSettings" tags:"settings" method:"get" summary:"系统设置"`
}

type ShowSettingsRes struct {
	g.Meta `mime:"json"`
	*model.ShowSettings
}

type SaveSettingsReq struct {
	g.Meta `path:"/save" tags:"settings" method:"post" summary:"保存员工信息"`
	*model.SaveSettings
}

type SaveSettingsRes struct {
	g.Meta `mime:"json"`
	Saved  int `json:"saved"`
}

type SettingsDetailReq struct {
	g.Meta `path:"/userDetail" tags:"settings" method:"get" summary:"显示员工设置"`
}

type SettingsDetailRes struct {
	g.Meta `mime:"json"`
	*model.UserDetail
}
