package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
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
