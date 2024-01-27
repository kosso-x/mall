// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsExceptAreaDetail is the golang structure of table hiolabs_except_area_detail for DAO operations like Where/Data.
type HiolabsExceptAreaDetail struct {
	g.Meta       `orm:"table:hiolabs_except_area_detail, do:true"`
	Id           interface{} //
	ExceptAreaId interface{} //
	Area         interface{} // 0位默认，
	IsDelete     interface{} //
}
