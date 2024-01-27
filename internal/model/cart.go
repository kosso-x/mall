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
	Id                           int
	User_id                      int
	Goods_id                     int
	Goods_sn                     int
	Product_id                   int
	Goods_name                   string
	Goods_aka                    string
	Goods_weight                 int
	Add_price                    float32
	Retail_price                 float32
	Number                       int
	Goods_specifition_name_value string
	Goods_specifition_ids        int
	Checked                      int
	List_pic_url                 string
	Freight_template_id          int
	Is_on_sale                   int
	Add_time                     int
	Is_fast                      int
	Is_delete                    int
	Weight_count                 int
}

type CartTotal struct {
	GoodsCount         int
	GoodsAmount        float32
	CheckedGoodsCount  int
	CheckedGoodsAmount float32
	User_id            int
	NumberChange       int
}

// update
type UpdateCart struct {
	Id        int
	Number    int
	ProductId int
}

// checked
type CheckedCart struct {
	ProductIds []int
	IsChecked  int
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
	CheckedAddress   int
	FreightPrice     float32
	CheckedGoodsList []CartList
	GoodsTotalPrice  float32
	OrderTotalPrice  float32
	ActualPrice      float32
	GoodsCount       int
	OutStock         int
	NumberChange     int
}
