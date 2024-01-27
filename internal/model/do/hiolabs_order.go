// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrder is the golang structure of table hiolabs_order for DAO operations like Where/Data.
type HiolabsOrder struct {
	g.Meta         `orm:"table:hiolabs_order, do:true"`
	Id             interface{} //
	OrderSn        interface{} //
	UserId         interface{} //
	OrderStatus    interface{} // 101：未付款、102：已取消、103已取消(系统)、201：已付款、202：订单取消，退款中、203：已退款、301：已发货、302：已收货、303：已收货(系统)、401：已完成、801：拼团中,未付款、802：拼团中,已付款
	OfflinePay     interface{} // 线下支付订单标志
	ShippingStatus interface{} // 0未发货，1已发货
	PrintStatus    interface{} //
	PayStatus      interface{} //
	Consignee      interface{} //
	Country        interface{} //
	Province       interface{} //
	City           interface{} //
	District       interface{} //
	Address        interface{} //
	PrintInfo      interface{} //
	Mobile         interface{} //
	Postscript     interface{} //
	AdminMemo      interface{} //
	ShippingFee    interface{} // 免邮的商品的邮费，这个在退款时不能退给客户
	PayName        interface{} //
	PayId          interface{} //
	ChangePrice    interface{} // 0没改价，不等于0改过价格，这里记录原始的价格
	ActualPrice    interface{} // 实际需要支付的金额
	OrderPrice     interface{} // 订单总价
	GoodsPrice     interface{} // 商品总价
	AddTime        interface{} //
	PayTime        interface{} // 付款时间
	ShippingTime   interface{} // 发货时间
	ConfirmTime    interface{} // 确认时间
	DealdoneTime   interface{} // 成交时间，用户评论或自动好评时间
	FreightPrice   interface{} // 配送费用
	ExpressValue   interface{} // 顺丰保价金额
	Remark         interface{} //
	OrderType      interface{} // 订单类型：0普通，1秒杀，2团购，3返现订单,7充值，8会员
	IsDelete       interface{} // 订单删除标志
}
