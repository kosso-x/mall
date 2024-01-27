// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsProduct is the golang structure for table hiolabs_product.
type HiolabsProduct struct {
	Id                    uint    `json:"id"                    ` //
	GoodsId               uint    `json:"goodsId"               ` //
	GoodsSpecificationIds string  `json:"goodsSpecificationIds" ` //
	GoodsSn               string  `json:"goodsSn"               ` //
	GoodsNumber           uint    `json:"goodsNumber"           ` //
	RetailPrice           float64 `json:"retailPrice"           ` //
	Cost                  float64 `json:"cost"                  ` // 成本
	GoodsWeight           float64 `json:"goodsWeight"           ` // 重量
	HasChange             int     `json:"hasChange"             ` //
	GoodsName             string  `json:"goodsName"             ` //
	IsOnSale              int     `json:"isOnSale"              ` //
	IsDelete              int     `json:"isDelete"              ` //
}
