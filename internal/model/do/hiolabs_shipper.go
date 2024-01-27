// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsShipper is the golang structure of table hiolabs_shipper for DAO operations like Where/Data.
type HiolabsShipper struct {
	g.Meta       `orm:"table:hiolabs_shipper, do:true"`
	Id           interface{} //
	Name         interface{} // 快递公司名称
	Code         interface{} // 快递公司代码
	SortOrder    interface{} // 排序
	MonthCode    interface{} //
	CustomerName interface{} //
	Enabled      interface{} //
}
