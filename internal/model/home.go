package model

// index/appInfo
type Info struct {
	Channel      []ShortCategory
	Banner       []ShortAd
	Notice       []ShortNotice
	CategoryList []CategoryGood
	CartCount    int
}

type ShortCategory struct {
	Id         int    `json:"id"`
	Icon_url   string `json:"icon_url"`
	Name       string `json:"name"`
	Sort_order int    `json:"sort_order"`
}

type ShortAd struct {
	Link_type int    `json:"link_type"`
	Goods_id  int    `json:"goods_id"`
	Image_url string `json:"image_url"`
	Link      string `json:"link"`
}

type ShortNotice struct {
	Content string `json:"content"`
}

type CategoryGood struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Banner    string `json:"banner"`
	Height    int    `json:"height"`
	GoodsList []GoodsList
}

type GoodsList struct {
	Id               int     `json:"id"`
	List_pic_url     string  `json:"list_pic_url"`
	Is_new           int     `json:"is_new"`
	Goods_number     int     `json:"goods_number"`
	Name             string  `json:"name"`
	Min_retail_price float32 `json:"min_retail_price"`
}
