// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrderExpress is the golang structure of table hiolabs_order_express for DAO operations like Where/Data.
type HiolabsOrderExpress struct {
	g.Meta       `orm:"table:hiolabs_order_express, do:true"`
	Id           interface{} //
	OrderId      interface{} //
	ShipperId    interface{} //
	ShipperName  interface{} // 物流公司名称
	ShipperCode  interface{} // 物流公司代码
	LogisticCode interface{} // 快递单号
	Traces       interface{} // 物流跟踪信息
	IsFinish     interface{} //
	RequestCount interface{} // 总查询次数
	RequestTime  interface{} // 最近一次向第三方查询物流信息时间
	AddTime      interface{} // 添加时间
	UpdateTime   interface{} // 更新时间
	ExpressType  interface{} //
	RegionCode   interface{} // 快递的地区编码，如杭州571
}
