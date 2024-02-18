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

type GetBase64Req struct {
	g.Meta  `path:"/qrcode/getBase64" tags:"qrcode" method:"post" summary:"获取商品详情二维码"`
	GoodsId int
}

type GetBase64Res struct {
	g.Meta     `mime:"json"`
	Ciphertext string `json:"ciphertext"`
}
