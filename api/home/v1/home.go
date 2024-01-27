package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AppInfoReq struct {
	g.Meta `path:"/index/appInfo" tags:"app-info" method:"get" summary:"首页数据接口"`
}
type AppInfoRes struct {
	g.Meta       `mime:"json"`
	Channel      *[]model.ShortCategory `json:"channel"`
	Banner       *[]model.ShortAd       `json:"banner"`
	Notice       *[]model.ShortNotice   `json:"notice"`
	CategoryList *[]model.CategoryGood  `json:"categoryList"`
	CartCount    *int                   `json:"cartCount"`
}
