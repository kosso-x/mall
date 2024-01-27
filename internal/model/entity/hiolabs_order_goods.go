// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsOrderGoods is the golang structure for table hiolabs_order_goods.
type HiolabsOrderGoods struct {
	Id                        uint    `json:"id"                        ` //
	OrderId                   uint    `json:"orderId"                   ` //
	GoodsId                   uint    `json:"goodsId"                   ` //
	GoodsName                 string  `json:"goodsName"                 ` //
	GoodsAka                  string  `json:"goodsAka"                  ` //
	ProductId                 uint    `json:"productId"                 ` //
	Number                    uint    `json:"number"                    ` //
	RetailPrice               float64 `json:"retailPrice"               ` //
	GoodsSpecifitionNameValue string  `json:"goodsSpecifitionNameValue" ` //
	GoodsSpecifitionIds       string  `json:"goodsSpecifitionIds"       ` //
	ListPicUrl                string  `json:"listPicUrl"                ` //
	UserId                    uint    `json:"userId"                    ` //
	IsDelete                  int     `json:"isDelete"                  ` // 是否删除标志
}
