// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsGoods is the golang structure for table hiolabs_goods.
type HiolabsGoods struct {
	Id                uint    `json:"id"                ` //
	CategoryId        uint    `json:"categoryId"        ` //
	IsOnSale          uint    `json:"isOnSale"          ` //
	Name              string  `json:"name"              ` //
	GoodsNumber       uint    `json:"goodsNumber"       ` //
	SellVolume        uint    `json:"sellVolume"        ` // 销售量
	Keywords          string  `json:"keywords"          ` //
	RetailPrice       string  `json:"retailPrice"       ` // 零售价格
	MinRetailPrice    float64 `json:"minRetailPrice"    ` //
	CostPrice         string  `json:"costPrice"         ` //
	MinCostPrice      float64 `json:"minCostPrice"      ` //
	GoodsBrief        string  `json:"goodsBrief"        ` //
	GoodsDesc         string  `json:"goodsDesc"         ` //
	SortOrder         uint    `json:"sortOrder"         ` //
	IsIndex           int     `json:"isIndex"           ` //
	IsNew             int     `json:"isNew"             ` //
	GoodsUnit         string  `json:"goodsUnit"         ` // 商品单位
	HttpsPicUrl       string  `json:"httpsPicUrl"       ` // 商品https图
	ListPicUrl        string  `json:"listPicUrl"        ` // 商品列表图
	FreightTemplateId int     `json:"freightTemplateId" ` //
	FreightType       int     `json:"freightType"       ` //
	IsDelete          uint    `json:"isDelete"          ` //
	HasGallery        int     `json:"hasGallery"        ` //
	HasDone           int     `json:"hasDone"           ` //
}
