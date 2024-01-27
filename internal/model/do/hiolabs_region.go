// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsRegion is the golang structure of table hiolabs_region for DAO operations like Where/Data.
type HiolabsRegion struct {
	g.Meta   `orm:"table:hiolabs_region, do:true"`
	Id       interface{} //
	ParentId interface{} //
	Name     interface{} //
	Type     interface{} //
	AgencyId interface{} //
	Area     interface{} // 方位，根据这个定运费
	AreaCode interface{} // 方位代码
	FarArea  interface{} // 偏远地区
}
