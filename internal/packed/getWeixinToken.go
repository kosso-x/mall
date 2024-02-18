package packed

import (
	"context"
	"fmt"
	"mall/internal/model"

	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

func GetAuthToken(ctx context.Context) (res *model.WxToken, err error) {
	config := g.Cfg("weixin_auth")
	gctx := gctx.New()
	token_g_type, _ := config.Get(gctx, "weixin.token_g_type")
	secret, _ := config.Get(gctx, "weixin.secret")
	appid, _ := config.Get(gctx, "weixin.appid")
	genurl, _ := config.Get(gctx, "weixin.token_url")

	url := fmt.Sprintf("%s?grant_type=%s&secret=%s&appid=%s", genurl, token_g_type, secret, appid)
	resp, err := g.Client().Get(gctx, url)

	if err != nil {
		glog.Errorf(gctx, "request %v error", url)
		return
	}

	defer resp.Close()
	body_json := resp.ReadAllString()

	if j, err := gjson.DecodeToJson(body_json); err != nil {
		return nil, err
	} else {
		if err := j.Scan(&res); err != nil {
			return nil, err
		}
		if res.Errmsg != "" {
			return nil, gerror.New(res.Errmsg)
		}
	}

	return
}

func GetImageUrl(ctx context.Context, token_res *model.WxToken, goodsId int) (image_url string, err error) {
	genurl, _ := g.Cfg("weixin_auth").Get(ctx, "weixin.wx_code_url")
	url := fmt.Sprintf("%s?access_token=%s", genurl, token_res.Access_token)

	resp := g.Client().ContentJson().PostContent(ctx, url, g.Map{
		"scene":      goodsId,
		"page":       "pages/goods/goods",
		"width":      200,
		"check_path": false,
	})
	image_url = gbase64.EncodeString(resp)

	return
}
