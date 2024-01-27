// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFormid is the golang structure of table hiolabs_formid for DAO operations like Where/Data.
type HiolabsFormid struct {
	g.Meta   `orm:"table:hiolabs_formid, do:true"`
	Id       interface{} //
	UserId   interface{} //
	OrderId  interface{} //
	FormId   interface{} //
	AddTime  interface{} // 添加时间
	UseTimes interface{} // 使用次数
}
