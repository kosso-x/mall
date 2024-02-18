package model

// list
// req
type OrderListReq struct {
	ShowType int
	Size     int
	Page     int
}

// res
type OrderList struct {
	Count       int          `json:"count"`
	TotalPages  int          `json:"totalPages"`
	PageSize    int          `json:"pageSize"`
	CurrentPage int          `json:"currentPage"`
	Data        []OrderShort `json:"data"`
}

type OrderShort struct {
	Id              int          `json:"id"`
	AddTime         int          `json:"add_time"`
	ActualPrice     float32      `json:"actual_price"`
	FreightPrice    float32      `json:"freight_price"`
	OfflinePay      int          `json:"offline_pay"`
	OrderStatus     int          `json:"order_status"`
	GoodsList       []GoodsShort `json:"goodsList"`
	GoodsCount      int          `json:"goodsCount"`
	OrderStatusText string       `json:"order_status_text"`
	HandleOption    HandleOption `json:"handleOption"`
}

type GoodsShort struct {
	Id           int    `json:"id"`
	List_pic_url string `json:"list_pic_url"`
	Number       int    `json:"number"`
}

type HandleOption struct {
	Cancel       bool `json:"cancel"`
	Delete       bool `json:"delete"`
	Pay          bool `json:"pay"`
	Confirm      bool `json:"confirm"`
	CancelRefund bool `json:"cancel_refund"`
}

// submit
type OrderSubmit struct {
	AddressId    int     `json:"addressId"`
	Postscript   string  `json:"postscript"`
	FreightPrice int     `json:"freightPrice"`
	ActualPrice  float32 `json:"actualPrice"`
	OfflinePay   int     `json:"offlinePay"`
}

type OrderInfo struct {
	Actual_price  float32 `json:"actual_price"`
	Add_time      int     `json:"add_time"`
	Address       string  `json:"address"`
	Change_price  float32 `json:"change_price"`
	City          int     `json:"city"`
	Consignee     string  `json:"consignee"`
	District      int     `json:"district"`
	Freight_price float32 `json:"freight_price"`
	Goods_price   float32 `json:"goods_price"`
	Id            int     `json:"id"`
	Mobile        string  `json:"mobile"`
	Offline_pay   int     `json:"offline_pay"`
	Order_price   float32 `json:"order_price"`
	Order_sn      string  `json:"order_sn"`
	Order_status  int     `json:"order_status"`
	Postscript    string  `json:"postscript"`
	Print_info    string  `json:"print_info"`
	Province      int     `json:"province"`
	User_id       int     `json:"user_id"`
}

// detail
type OrderDetail struct {
	OrderInfo    OrderDetailInfo `json:"orderInfo"`
	OrderGoods   []HOrderGoods   `json:"orderGoods"`
	HandleOption HandleOption    `json:"handleOption"`
	TextCode     TextCode        `json:"textCode"`
	GoodsCount   int             `json:"goodsCount"`
}

type HOrderGoods struct {
	Id                        uint    `json:"id"                        `    //
	OrderId                   uint    `json:"order_id"                   `   //
	GoodsId                   uint    `json:"goods_id"                   `   //
	GoodsName                 string  `json:"goods_name"                 `   //
	GoodsAka                  string  `json:"goods_aka"                  `   //
	ProductId                 uint    `json:"product_id"                 `   //
	Number                    uint    `json:"number"                    `    //
	RetailPrice               float64 `json:"retail_price"               `   //
	GoodsSpecifitionNameValue string  `json:"goods_specifition_name_value" ` //
	GoodsSpecifitionIds       string  `json:"goods_specifition_ids"       `  //
	ListPicUrl                string  `json:"list_pic_url"                `  //
	UserId                    uint    `json:"user_id"                    `   //
	IsDelete                  int     `json:"is_delete"                  `   // 是否删除标志
}

type OrderDetailInfo struct {
	HOrder
	RegionValues
	Order_status_text string `json:"order_status_text"`
	Final_pay_time    int    `json:"final_pay_time"`
}

type HOrder struct {
	Id             uint    `json:"id"             `  //
	OrderSn        string  `json:"order_sn"        ` //
	UserId         uint    `json:"user_id"         ` //
	OrderStatus    uint    `json:"order_status"    ` // 101：未付款、102：已取消、103已取消(系统)、201：已付款、202：订单取消，退款中、203：已退款、301：已发货、302：已收货、303：已收货(系统)、401：已完成、801：拼团中,未付款、802：拼团中,已付款
	OfflinePay     uint    `json:"offline_pay"     ` // 线下支付订单标志
	ShippingStatus uint    `json:"shipping_status" ` // 0未发货，1已发货
	PrintStatus    int     `json:"print_status"    ` //
	PayStatus      uint    `json:"pay_status"      ` //
	Consignee      string  `json:"consignee"      `  //
	Country        uint    `json:"country"        `  //
	Province       uint    `json:"province"       `  //
	City           uint    `json:"city"           `  //
	District       uint    `json:"district"       `  //
	Address        string  `json:"address"        `  //
	PrintInfo      string  `json:"print_info"      ` //
	Mobile         string  `json:"mobile"         `  //
	Postscript     string  `json:"postscript"     `  //
	AdminMemo      string  `json:"admin_memo"      ` //
	ShippingFee    float64 `json:"shipping_fee"    ` // 免邮的商品的邮费，这个在退款时不能退给客户
	PayName        string  `json:"pay_name"        ` //
	PayId          string  `json:"pay_id"          ` //
	ChangePrice    float64 `json:"change_price"    ` // 0没改价，不等于0改过价格，这里记录原始的价格
	ActualPrice    float64 `json:"actual_price"    ` // 实际需要支付的金额
	OrderPrice     float64 `json:"order_price"     ` // 订单总价
	GoodsPrice     float64 `json:"goods_price"     ` // 商品总价
	AddTime        uint    `json:"add_time"        ` //
	PayTime        uint    `json:"pay_time"        ` // 付款时间
	ShippingTime   uint    `json:"shipping_time"   ` // 发货时间
	ConfirmTime    uint    `json:"confirm_time"    ` // 确认时间
	DealdoneTime   uint    `json:"dealdone_time"   ` // 成交时间，用户评论或自动好评时间
	FreightPrice   uint    `json:"freight_price"   ` // 配送费用
	ExpressValue   float64 `json:"express_value"   ` // 顺丰保价金额
	Remark         string  `json:"remark"         `  //
	OrderType      uint    `json:"order_type"      ` // 订单类型：0普通，1秒杀，2团购，3返现订单,7充值，8会员
	IsDelete       uint    `json:"is_delete"       ` // 订单删除标志
}

type TextCode struct {
	Pay       bool `json:"pay"`
	Close     bool `json:"close"`
	Delivery  bool `json:"delivery"`
	Receive   bool `json:"receive"`
	Success   bool `json:"success"`
	Countdown bool `json:"countdown"`
}

// orderCount
type OrderCountInfo struct {
	ToPay      int `json:"toPay"`
	ToDelivery int `json:"toDelivery"`
	ToReceive  int `json:"toReceive"`
}

// OrderGoods
type (
	OrderGoodsRes []HOrderGoods
)
