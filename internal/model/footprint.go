package model

// list
type FootprintList struct {
	Count       int                      `json:"count"`
	TotalPages  int                      `json:"totalPages"`
	PageSize    int                      `json:"pageSize"`
	CurrentPage int                      `json:"currentPage"`
	Data        []FootprintInfoWithGoods `json:"data"`
}

type FootprintInfoWithGoods struct {
	Id      int            `json:"id"`
	GoodsId int            `json:"goods_id"`
	AddTime string         `json:"add_time"`
	Goods   FootprintGoods `json:"goods"`
}

type FootprintGoods struct {
	Name           string  `json:"name"`
	GoodsBrief     string  `json:"goods_brief"`
	RetailPrice    float32 `json:"retail_price"`
	ListPicUrl     string  `json:"list_pic_url"`
	GoodsNumber    int     `json:"goods_number"`
	MinRetailPrice float32 `json:"min_retail_price"`
}

type FootprintGoodsInfo struct {
	Id             int     `json:"id"`
	GoodsId        int     `json:"goods_id"`
	AddTime        string  `json:"add_time"`
	Name           string  `json:"name"`
	GoodsBrief     string  `json:"goods_brief"`
	RetailPrice    float32 `json:"retail_price"`
	ListPicUrl     string  `json:"list_pic_url"`
	GoodsNumber    int     `json:"goods_number"`
	MinRetailPrice float32 `json:"min_retail_price"`
}

// delete
