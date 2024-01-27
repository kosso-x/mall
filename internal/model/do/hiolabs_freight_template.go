// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplate is the golang structure of table hiolabs_freight_template for DAO operations like Where/Data.
type HiolabsFreightTemplate struct {
	g.Meta       `orm:"table:hiolabs_freight_template, do:true"`
	Id           interface{} //
	Name         interface{} // 运费模板名称
	PackagePrice interface{} // 包装费用
	FreightType  interface{} // 0按件，1按重量
	IsDelete     interface{} //
}
