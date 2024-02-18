package model

// add
type CartAddReq struct {
	AddType   int
	GoodsId   string
	Number    int
	ProductId int
}

type CartAddRes struct {
	CartList  []CartList
	CartTotal CartTotal
}

type CartList struct {
	Id                           int     `json:"id"`
	User_id                      int     `json:"user_id"`
	Goods_id                     int     `json:"goods_id"`
	Goods_sn                     int     `json:"goods_sn"`
	Product_id                   int     `json:"product_id"`
	Goods_name                   string  `json:"goods_name"`
	Goods_aka                    string  `json:"goods_aka"`
	Goods_weight                 int     `json:"goods_weight"`
	Add_price                    float32 `json:"add_price"`
	Retail_price                 float32 `json:"retail_price"`
	Number                       int     `json:"number"`
	Goods_specifition_name_value string  `json:"goods_specifition_name_value"`
	Goods_specifition_ids        int     `json:"goods_specifition_ids"`
	Checked                      int     `json:"checked"`
	List_pic_url                 string  `json:"list_pic_url"`
	Freight_template_id          int     `json:"freight_template_id"`
	Is_on_sale                   int     `json:"is_on_sale"`
	Add_time                     int     `json:"add_time"`
	Is_fast                      int     `json:"is_fast"`
	Is_delete                    int     `json:"is_delete"`
	Weight_count                 int     `json:"weight_count"`
}

type CartTotal struct {
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float32 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float32 `json:"checkedGoodsAmount"`
	User_id            int     `json:"user_id"`
	NumberChange       int     `json:"numberChange"`
}

// update
type UpdateCart struct {
	Id        int `json:"id"`
	Number    int `json:"number"`
	ProductId int `json:"productId"`
}

// checked
type CheckedCart struct {
	ProductIds []int `json:"productIds"`
	IsChecked  int   `json:"isChecked"`
}

// goodsCount
type CartGoodsCount struct {
	CartTotal GoodsCount
}

type GoodsCount struct {
	GoodsCount int
}

// checkout
type CartCheckoutReq struct {
	AddressId int
	AddType   int
	OrderFrom string
	Type      int
}

type CartCheckout struct {
	CheckedAddress   CartAddressInfo `json:"checkedAddress"`
	FreightPrice     float32         `json:"freightPrice"`
	CheckedGoodsList []CartList      `json:"checkedGoodsList"`
	GoodsTotalPrice  float32         `json:"goodsTotalPrice"`
	OrderTotalPrice  float32         `json:"orderTotalPrice"`
	ActualPrice      float32         `json:"actualPrice"`
	GoodsCount       int             `json:"goodsCount"`
	OutStock         int             `json:"outStock"`
	NumberChange     int             `json:"numberChange"`
}

type CartAddressInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	User_id     int    `json:"user_id"`
	Country_id  int    `json:"country_id"`
	Province_id int    `json:"province_id"`
	City_id     int    `json:"city_id"`
	District_id int    `json:"district_id"`
	Address     string `json:"address"`
	Mobile      string `json:"mobile"`
	Is_default  int    `json:"is_default"`
	Is_delete   int    `json:"is_delete"`
	RegionValues
}
