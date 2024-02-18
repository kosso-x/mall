package cmd

import (
	"fmt"
	"mall/internal/dao"
	"mall/internal/model"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/golang-jwt/jwt/v5"
)

// 更新token缓存中间件
func UpdateTokenHandle(r *ghttp.Request) {
	user, _, err := UnmarshalToken(r)

	if err != nil {
		return
	}

	r.Middleware.Next()
	/*
		1. 获取token
		2. 解析出结果
		3. 更新缓存 (user_info, token)
	*/
	expire, _ := g.Cfg().Get(ctx, "redis.default.expire")

	g.Redis().Do(ctx, "expire", fmt.Sprintf("mall-token-%s", user.Openid), expire)
	g.Redis().Do(ctx, "expire", fmt.Sprintf("mall-userid-%s", user.Openid), expire)
}

func UnmarshalToken(r *ghttp.Request) (user *model.SignTokenUser, tokenString string, err error) {
	tokenString = r.Header.Get("X-Hioshop-Token")
	secret, _ := g.Cfg("weixin_auth").Get(ctx, "weixin.secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			glog.Errorf(ctx, "Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(gconv.Bytes(secret)), nil
	})

	if err != nil {
		glog.Error(ctx, err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user = &model.SignTokenUser{
			Session_key: gconv.String(claims["session_key"]),
			Openid:      gconv.String(claims["openid"]),
			User_id:     gconv.Int(claims["user_id"]),
			Iat:         gconv.Int(claims["iat"]),
		}
	} else {
		glog.Error(ctx, err)
	}

	dao.HiolabsUser.Ctx(ctx).Where("id = ?", user.User_id).Scan(&CurrentUser)
	return
}
