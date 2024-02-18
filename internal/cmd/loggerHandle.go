package cmd

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func LoggerHandle(r *ghttp.Request) {
	json, _ := r.GetJson()
	glog.Info(ctx, r.Method, fmt.Sprintf("%s%s", r.Host, r.RequestURI), r.Proto, json)

	r.Middleware.Next()
}
