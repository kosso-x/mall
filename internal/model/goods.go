package model

// count

// detail
type GoodsDetail struct {
	Info              HGoods             `json:"info"`
	Gallery           []HGoodsGallery    `json:"gallery"`
	SpecificationList GoodsSpecification `json:"specificationList"`
	ProductList       []HProduct         `json:"productList"`
}

type HGoods struct {
	Id                uint    `json:"id"                `   //
	CategoryId        uint    `json:"category_id"        `  //
	IsOnSale          uint    `json:"is_on_sale"          ` //
	Name              string  `json:"name"              `   //
	GoodsNumber       uint    `json:"goods_number"       `  //
	SellVolume        uint    `json:"sell_volume"        `  // 销售量
	Keywords          string  `json:"keywords"          `   //
	RetailPrice       string  `json:"retail_price"       `  // 零售价格
	MinRetailPrice    float64 `json:"min_retail_price"    ` //
	CostPrice         string  `json:"cost_price"         `  //
	MinCostPrice      float64 `json:"min_cost_price"      ` //
	GoodsBrief        string  `json:"goods_brief"        `  //
	GoodsDesc         string  `json:"goods_desc"         `  //
	SortOrder         uint    `json:"sort_order"         `  //
	IsIndex           int     `json:"is_index"           `  //
	IsNew             int     `json:"is_new"             `  //
	GoodsUnit         string  `json:"goods_unit"         `  // 商品单位
	HttpsPicUrl       string  `json:"https_pic_url"       ` // 商品https图
	ListPicUrl        string  `json:"list_pic_url"        ` // 商品列表图
	FreightTemplateId int     `json:"freight_template_id" ` //
	FreightType       int     `json:"freight_type"       `  //
	IsDelete          uint    `json:"is_delete"          `  //
	HasGallery        int     `json:"has_gallery"        `  //
	HasDone           int     `json:"has_done"           `  //
}

type HGoodsGallery struct {
	Id        uint   `json:"id"        `  //
	GoodsId   uint   `json:"goods_id"   ` //
	ImgUrl    string `json:"img_url"    ` //
	ImgDesc   string `json:"img_desc"   ` //
	SortOrder uint   `json:"sort_order" ` //
	IsDelete  int    `json:"is_delete"  ` //
}

type HProduct struct {
	Id                    uint    `json:"id"                    `   //
	GoodsId               uint    `json:"goods_id"               `  //
	GoodsSpecificationIds string  `json:"goods_specification_ids" ` //
	GoodsSn               string  `json:"goods_sn"               `  //
	GoodsNumber           uint    `json:"goods_number"           `  //
	RetailPrice           float64 `json:"retail_price"           `  //
	Cost                  float64 `json:"cost"                  `   // 成本
	GoodsWeight           float64 `json:"goods_weight"           `  // 重量
	HasChange             int     `json:"has_change"             `  //
	GoodsName             string  `json:"goods_name"             `  //
	IsOnSale              int     `json:"is_on_sale"              ` //
	IsDelete              int     `json:"is_delete"              `  //
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
	GoodsListInfo []HGoods
)

// goodsShare

type GoodsShare struct {
	Name         string  `json:"name"`
	Retail_price float32 `json:"retail_price"`
}
