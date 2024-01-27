package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AuthLoginReq struct {
	g.Meta `path:"/loginByWeixin" tags:"login" method:"post" summary:"微信登录"`
	Code   string
}
type AuthLoginRes struct {
	g.Meta `mime:"json"`
	*model.AuthLogin
}
