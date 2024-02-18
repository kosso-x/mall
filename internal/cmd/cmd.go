package cmd

import (
	"context"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"

	"mall/internal/controller/address"
	"mall/internal/controller/cart"
	"mall/internal/controller/category"
	"mall/internal/controller/footprint"
	"mall/internal/controller/goods"
	"mall/internal/controller/home"
	"mall/internal/controller/login"
	"mall/internal/controller/order"
	"mall/internal/controller/region"
	"mall/internal/controller/search"
	"mall/internal/controller/settings"
	"mall/internal/controller/upload"
	"mall/internal/model"
)

var (
	ctx         = gctx.New()
	CurrentUser *model.LoginUserInfo
	Main        = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			SetDefaultLogger() // 替代默认的log
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ResponseHandler)
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
				group.Group("/region", func(group *ghttp.RouterGroup) {
					group.Bind(
						region.NewV1(),
					)
				})
				group.Group("/order", func(group *ghttp.RouterGroup) {
					group.Bind(
						order.NewV1(),
					)
				})
				group.Group("/footprint", func(group *ghttp.RouterGroup) {
					group.Bind(
						footprint.NewV1(),
					)
				})
				group.Group("/search", func(group *ghttp.RouterGroup) {
					group.Bind(
						search.NewV1(),
					)
				})
				group.Group("/settings", func(group *ghttp.RouterGroup) {
					group.Bind(
						settings.NewV1(),
					)
				})
				group.Group("/upload", func(group *ghttp.RouterGroup) {
					group.Bind(
						upload.NewV1(),
					)
				})

				group.Bind(
					home.NewV1(),
				)
			})

			s.SetIndexFolder(true)
			static_path := g.Cfg("mall").MustGetWithCmd(ctx, "default.static_path")
			s.SetServerRoot(gconv.String(static_path))
			s.Run()
			return nil
		},
	}
)
