package main

import (
	_ "mall/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/os/gctx"

	"mall/internal/cmd"
	_ "mall/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
