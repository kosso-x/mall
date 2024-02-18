package model

import "mall/internal/model/entity"

// catalog/index
type (
	CategoryList []entity.HiolabsCategory
)

// catalog/current
type CurrentCategory struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Img_url  string `json:"img_url"`
	P_height int    `json:"p_height"`
}

// catalog/currentlist
type CategoryCurrentList struct {
	Count       int          `json:"count"`
	TotalPages  int          `json:"totalPages"`
	PageSize    int          `json:"pageSize"`
	CurrentPage int          `json:"currentPage"`
	Data        []ShortGoods `json:"data"`
}

type ShortGoods struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Goods_brief      string  `json:"goods_brief"`
	Min_retail_price float32 `json:"min_retail_price"`
	List_pic_url     string  `json:"list_pic_url"`
	Goods_number     int     `json:"goods_number"`
}
