package settings

import (
	"context"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/service"

	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sSettings struct{}
)

func init() {
	service.RegisterSettings(New())
}

func New() service.ISettings {
	return &sSettings{}
}

func (s *sSettings) Show(ctx context.Context) (res *model.ShowSettings, err error) {
	err = dao.HiolabsShowSettings.Ctx(ctx).Where("id = ?", 1).Scan(&res)
	if err != nil {
		return
	}

	return
}

func (s *sSettings) Save(ctx context.Context, settings *model.SaveSettings) (err error) {
	name_mobile := 0
	if settings.Name != "" && settings.Mobile != "" {
		name_mobile = 1
	}
	nickname := gbase64.EncodeString(settings.NickName)
	_, err = dao.HiolabsUser.Ctx(ctx).Update(g.Map{
		"name":        settings.Name,
		"mobile":      settings.Mobile,
		"nickname":    nickname,
		"avatar":      settings.Avatar,
		"name_mobile": name_mobile,
	}, "id = ?", cmd.CurrentUser.Id)
	if err != nil {
		return gerror.New("update hiolabs_user error")
	}

	return
}

func (s *sSettings) UserDetail(ctx context.Context) (res *model.UserDetail, err error) {
	err = dao.HiolabsUser.Ctx(ctx).Where("id = ?", cmd.CurrentUser.Id).Scan(&res)
	if err != nil {
		return
	}
	res.Nickname, err = gbase64.DecodeToString(res.Nickname)
	if err != nil {
		return
	}
	return
}
