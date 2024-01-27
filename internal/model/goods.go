package model

import "mall/internal/model/entity"

// count

// detail
type GoodsDetail struct {
	Info              entity.HiolabsGoods          `json:"info"`
	Gallery           []entity.HiolabsGoodsGallery `json:"gallery"`
	SpecificationList GoodsSpecification           `json:"specificationList"`
	ProductList       []entity.HiolabsProduct      `json:"productList"`
}

type GoodsSpecification struct {
	Specification_id int         `json:"specification_id"`
	Name             string      `json:"name"`
	ValueList        []ValueList `json:"valueList"`
}

type ValueList struct {
	Id               int    `json:"id"`
	Goods_id         int    `json:"goods_id"`
	Specification_id int    `json:"specification_id"`
	Value            string `json:"value"`
	Pic_url          string `json:"pic_url"`
	Is_delete        int    `json:"is_delete"`
	Goods_number     int    `json:"goods_number"`
}

// list
type (
	GoodsListInfo []entity.HiolabsGoods
)

// goodsShare

type GoodsShare struct {
	Name         string  `json:"name"`
	Retail_price float32 `json:"retail_price"`
}
