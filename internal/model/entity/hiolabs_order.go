// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsOrder is the golang structure for table hiolabs_order.
type HiolabsOrder struct {
	Id             uint    `json:"id"             ` //
	OrderSn        string  `json:"orderSn"        ` //
	UserId         uint    `json:"userId"         ` //
	OrderStatus    uint    `json:"orderStatus"    ` // 101：未付款、102：已取消、103已取消(系统)、201：已付款、202：订单取消，退款中、203：已退款、301：已发货、302：已收货、303：已收货(系统)、401：已完成、801：拼团中,未付款、802：拼团中,已付款
	OfflinePay     uint    `json:"offlinePay"     ` // 线下支付订单标志
	ShippingStatus uint    `json:"shippingStatus" ` // 0未发货，1已发货
	PrintStatus    int     `json:"printStatus"    ` //
	PayStatus      uint    `json:"payStatus"      ` //
	Consignee      string  `json:"consignee"      ` //
	Country        uint    `json:"country"        ` //
	Province       uint    `json:"province"       ` //
	City           uint    `json:"city"           ` //
	District       uint    `json:"district"       ` //
	Address        string  `json:"address"        ` //
	PrintInfo      string  `json:"printInfo"      ` //
	Mobile         string  `json:"mobile"         ` //
	Postscript     string  `json:"postscript"     ` //
	AdminMemo      string  `json:"adminMemo"      ` //
	ShippingFee    float64 `json:"shippingFee"    ` // 免邮的商品的邮费，这个在退款时不能退给客户
	PayName        string  `json:"payName"        ` //
	PayId          string  `json:"payId"          ` //
	ChangePrice    float64 `json:"changePrice"    ` // 0没改价，不等于0改过价格，这里记录原始的价格
	ActualPrice    float64 `json:"actualPrice"    ` // 实际需要支付的金额
	OrderPrice     float64 `json:"orderPrice"     ` // 订单总价
	GoodsPrice     float64 `json:"goodsPrice"     ` // 商品总价
	AddTime        uint    `json:"addTime"        ` //
	PayTime        uint    `json:"payTime"        ` // 付款时间
	ShippingTime   uint    `json:"shippingTime"   ` // 发货时间
	ConfirmTime    uint    `json:"confirmTime"    ` // 确认时间
	DealdoneTime   uint    `json:"dealdoneTime"   ` // 成交时间，用户评论或自动好评时间
	FreightPrice   uint    `json:"freightPrice"   ` // 配送费用
	ExpressValue   float64 `json:"expressValue"   ` // 顺丰保价金额
	Remark         string  `json:"remark"         ` //
	OrderType      uint    `json:"orderType"      ` // 订单类型：0普通，1秒杀，2团购，3返现订单,7充值，8会员
	IsDelete       uint    `json:"isDelete"       ` // 订单删除标志
}
