package service

import (
	"context"
	"mall/internal/model"
	"mall/internal/model/entity"
)

type (
	ILogin interface {
		AuthLogin(ctx context.Context, code string) (res *model.AuthLogin, err error)
		WeixinAuth(code string) (res *model.WxOpenAi, err error)
		UserRegister(ctx context.Context, user *entity.HiolabsUser, openid string) (err error)
		UpdateLogin(ctx context.Context, user *entity.HiolabsUser) (err error)
		GenerateToken(ctx context.Context, user *model.SignTokenUser) (token string, err error)
		StorageToken(ctx context.Context, user *model.SignTokenUser, token string) (err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
