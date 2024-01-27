// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsOrderExpress is the golang structure for table hiolabs_order_express.
type HiolabsOrderExpress struct {
	Id           uint   `json:"id"           ` //
	OrderId      uint   `json:"orderId"      ` //
	ShipperId    uint   `json:"shipperId"    ` //
	ShipperName  string `json:"shipperName"  ` // 物流公司名称
	ShipperCode  string `json:"shipperCode"  ` // 物流公司代码
	LogisticCode string `json:"logisticCode" ` // 快递单号
	Traces       string `json:"traces"       ` // 物流跟踪信息
	IsFinish     int    `json:"isFinish"     ` //
	RequestCount int    `json:"requestCount" ` // 总查询次数
	RequestTime  int    `json:"requestTime"  ` // 最近一次向第三方查询物流信息时间
	AddTime      int    `json:"addTime"      ` // 添加时间
	UpdateTime   int    `json:"updateTime"   ` // 更新时间
	ExpressType  int    `json:"expressType"  ` //
	RegionCode   string `json:"regionCode"   ` // 快递的地区编码，如杭州571
}
