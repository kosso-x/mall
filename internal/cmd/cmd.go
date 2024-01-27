package cmd

import (
	"context"

	"fmt"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"mall/internal/controller/address"
	"mall/internal/controller/cart"
	"mall/internal/controller/category"
	"mall/internal/controller/goods"
	"mall/internal/controller/home"
	"mall/internal/controller/login"
	"mall/internal/dao"
	"mall/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			SetDefaultLogger() // 替代默认的log
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(LoggerHandle)

				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						login.NewV1(),
					)
				})

				group.Middleware(UpdateTokenHandle)
				group.Group("/cart", func(group *ghttp.RouterGroup) {
					group.Bind(
						cart.NewV1(),
					)
				})
				group.Group("/catalog", func(group *ghttp.RouterGroup) {
					group.Bind(
						category.NewV1(),
					)
				})
				group.Group("/goods", func(group *ghttp.RouterGroup) {
					group.Bind(
						goods.NewV1(),
					)
				})
				group.Group("/address", func(group *ghttp.RouterGroup) {
					group.Bind(
						address.NewV1(),
					)
				})

				group.Bind(
					home.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
	ctx         = gctx.New()
	CurrentUser *model.LoginUserInfo
)

// 替代默认的日志handler，禁止控制台输出，全部输出到文件
func SetDefaultLogger() {
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout":   g.Config().MustGet(ctx, "logger.stdout"),
			"path":     g.Config().MustGet(ctx, "logger.path").String(), // 此处必须重新设置，才可以实现db的log写入文件
			"level":    g.Config().MustGet(ctx, "logger.level"),
			"StStatus": 0,
			"CtxKeys":  g.Config().MustGet(ctx, "logger.ctxKeys").String(),
		}
		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})
}

// 登录中间件
func LoginHandle(r *ghttp.Request) {
	r.Middleware.Next()
	/*
		1. 获取token
		2. 解析出结果
		3. (user_info, token)缓存到redis
	*/

}

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

func LoggerHandle(r *ghttp.Request) {
	json, _ := r.GetJson()
	glog.Info(ctx, r.Method, fmt.Sprintf("%s%s", r.Host, r.RequestURI), r.Proto, json)

	r.Middleware.Next()
}
