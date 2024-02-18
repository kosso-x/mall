package model

// show
type ShowSettings struct {
	Id               int `json:"id"`
	Banner           int `json:"banner"`
	Channel          int `json:"channel"`
	Index_banner_img int `json:"index_banner_img"`
	Notice           int `json:"notice"`
}

// save
type SaveSettings struct {
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

// userDetail
type UserDetail struct {
	Id       int    `json:"id"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
