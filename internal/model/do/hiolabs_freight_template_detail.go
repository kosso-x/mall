// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplateDetail is the golang structure of table hiolabs_freight_template_detail for DAO operations like Where/Data.
type HiolabsFreightTemplateDetail struct {
	g.Meta     `orm:"table:hiolabs_freight_template_detail, do:true"`
	Id         interface{} //
	TemplateId interface{} //
	GroupId    interface{} //
	Area       interface{} // 0位默认，
	IsDelete   interface{} //
}
