// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsCart is the golang structure for table hiolabs_cart.
type HiolabsCart struct {
	Id                        uint    `json:"id"                        ` //
	UserId                    uint    `json:"userId"                    ` //
	GoodsId                   uint    `json:"goodsId"                   ` //
	GoodsSn                   string  `json:"goodsSn"                   ` //
	ProductId                 uint    `json:"productId"                 ` //
	GoodsName                 string  `json:"goodsName"                 ` //
	GoodsAka                  string  `json:"goodsAka"                  ` //
	GoodsWeight               float64 `json:"goodsWeight"               ` // 重量
	AddPrice                  float64 `json:"addPrice"                  ` // 加入购物车时的价格
	RetailPrice               float64 `json:"retailPrice"               ` //
	Number                    uint    `json:"number"                    ` //
	GoodsSpecifitionNameValue string  `json:"goodsSpecifitionNameValue" ` // 规格属性组成的字符串，用来显示用
	GoodsSpecifitionIds       string  `json:"goodsSpecifitionIds"       ` // product表对应的goods_specifition_ids
	Checked                   uint    `json:"checked"                   ` //
	ListPicUrl                string  `json:"listPicUrl"                ` //
	FreightTemplateId         uint    `json:"freightTemplateId"         ` // 运费模板
	IsOnSale                  int     `json:"isOnSale"                  ` // 0
	AddTime                   int     `json:"addTime"                   ` //
	IsFast                    int     `json:"isFast"                    ` // 1
	IsDelete                  uint    `json:"isDelete"                  ` //
}
