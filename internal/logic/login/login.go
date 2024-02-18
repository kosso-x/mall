package login

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gbase64"

	"fmt"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/model/entity"
	"mall/internal/packed"
	"mall/internal/service"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	sLogin struct{}
)

func init() {
	service.RegisterLogin(New())
}

func New() service.ILogin {
	return &sLogin{}
}

func (s *sLogin) AuthLogin(ctx context.Context, code string) (res *model.AuthLogin, err error) {
	// 获取openid
	wx_resp, err := s.WeixinAuth(code)

	if err != nil {
		return nil, gerror.New("登录失败，openid无效")
	}

	if wx_resp.Errcode != 0 {
		return nil, gerror.New(wx_resp.Errmsg)
	}

	// 看是否注册
	var h_user *entity.HiolabsUser
	// 注册的标记
	Is_new := 0
	dao.HiolabsUser.Ctx(ctx).Where("weixin_openid = ?", wx_resp.Openid).Scan(&h_user)
	if (h_user == nil) || (err != nil) {
		// 注册
		s.UserRegister(ctx, h_user, wx_resp.Openid)
		Is_new = 1
	}
	// 更新登录信息
	s.UpdateLogin(ctx, h_user)

	// 微信远程登录 token
	packed.GetAuthToken(ctx)
	// 微信本地登录token

	sing_user := &model.SignTokenUser{
		Session_key: wx_resp.Session_key,
		Openid:      wx_resp.Openid,
		User_id:     int(h_user.Id),
		Iat:         int(time.Now().Unix()),
	}
	tokenString, err := s.GenerateToken(ctx, sing_user)
	if err != nil {
		return
	}

	nickname, err := gbase64.DecodeToString(h_user.Nickname)
	if err != nil {
		return
	}

	res = &model.AuthLogin{
		Token: tokenString,
		UserInfo: model.LoginUserInfo{
			Id:       int(h_user.Id),
			Username: h_user.Username,
			Nickname: nickname,
			Avatar:   h_user.Avatar,
		},
		Is_new: Is_new,
	}

	login_user := &model.SignTokenUser{
		Session_key: wx_resp.Session_key,
		Openid:      wx_resp.Openid,
		User_id:     int(h_user.Id),
	}
	s.StorageToken(ctx, login_user, tokenString)

	return
}

func (s *sLogin) WeixinAuth(code string) (res *model.WxOpenAi, err error) {
	config := g.Cfg("weixin_auth")
	gctx := gctx.New()
	openai_g_type, _ := config.Get(gctx, "weixin.openai_g_type")
	secret, _ := config.Get(gctx, "weixin.secret")
	appid, _ := config.Get(gctx, "weixin.appid")
	genurl, _ := config.Get(gctx, "weixin.auth_url")

	url := fmt.Sprintf("%s?grant_type=%s&js_code=%s&secret=%s&appid=%s", genurl, openai_g_type, code, secret, appid)
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

func (s *sLogin) UserRegister(ctx context.Context, user *entity.HiolabsUser, openid string) (err error) {
	// 生成 微信用户 的 base64
	base_name := "微信用户"
	nickname := gbase64.EncodeString(base_name)
	// 生成uuid6
	username := base_name + uuid.New().String()
	// 当前时间戳
	current_time := time.Now().Unix()
	_, err = dao.HiolabsUser.Ctx(ctx).Insert(
		g.Map{
			"username":        username,
			"password":        openid,
			"register_time":   string(current_time),
			"register_ip":     "",
			"last_login_time": string(current_time),
			"last_login_ip":   "",
			"mobile":          "",
			"weixin_openid":   openid,
			"nickname":        nickname,
			"avatar":          "mall/resource/public/resource/image/default_avatar.png",
		})

	if err != nil {
		return gerror.New("登录失败")
	}

	dao.HiolabsUser.Ctx(ctx).Where("weixin_openid = ?", openid).Scan(&user)
	return
}

func (s *sLogin) UpdateLogin(ctx context.Context, user *entity.HiolabsUser) (err error) {
	current_time := time.Now().Unix()
	_, err = dao.HiolabsUser.Ctx(ctx).Update(
		g.Map{
			"last_login_time": current_time,
			"last_login_ip":   "",
		}, "id", user.Id)

	if err != nil {
		return gerror.New("更新登录信息失败")
	}

	return
}

func (s *sLogin) GenerateToken(ctx context.Context, user *model.SignTokenUser) (token string, err error) {
	gen_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"session_key": user.Session_key,
		"openid":      user.Openid,
		"user_id":     user.User_id,
		"iat":         user.Iat,
	})
	secret, _ := g.Cfg("weixin_auth").Get(ctx, "weixin.secret")

	token, err = gen_token.SignedString(gconv.Bytes(secret))
	if err != nil {
		return "", gerror.New("sign token error")
	}

	return
}

func (s *sLogin) StorageToken(ctx context.Context, user *model.SignTokenUser, token string) (err error) {
	expire, _ := g.Cfg().Get(ctx, "redis.default.expire")
	g.Redis().Do(ctx, "SET", fmt.Sprintf("mall-token-%s", user.Openid), token, "ex", expire)
	g.Redis().Do(ctx, "SET", fmt.Sprintf("mall-userid-%s", user.Openid), user.User_id, "ex", expire)
	return
}
