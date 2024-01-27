// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplateGroup is the golang structure of table hiolabs_freight_template_group for DAO operations like Where/Data.
type HiolabsFreightTemplateGroup struct {
	g.Meta       `orm:"table:hiolabs_freight_template_group, do:true"`
	Id           interface{} //
	TemplateId   interface{} //
	IsDefault    interface{} // 默认，area:0
	Area         interface{} // 0位默认，
	Start        interface{} //
	StartFee     interface{} //
	Add          interface{} //
	AddFee       interface{} //
	FreeByNumber interface{} // 0没有设置
	FreeByMoney  interface{} // 0没设置
	IsDelete     interface{} //
}
