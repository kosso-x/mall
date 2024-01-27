package model

type WxOpenAi struct {
	Session_key string
	Unionid     string
	Errmsg      string
	Openid      string
	Errcode     int32
}

type WxToken struct {
	Access_token string
	Expires_in   int
}

type AuthLogin struct {
	Token    string        `json:"token"`
	UserInfo LoginUserInfo `json:"userInfo"`
	Is_new   int           `json:"is_new"`
}

type LoginUserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// 注册token用
type SignTokenUser struct {
	Session_key string
	Openid      string
	User_id     int
	Iat         int
}
