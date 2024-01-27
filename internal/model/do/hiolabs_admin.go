// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAdmin is the golang structure of table hiolabs_admin for DAO operations like Where/Data.
type HiolabsAdmin struct {
	g.Meta        `orm:"table:hiolabs_admin, do:true"`
	Id            interface{} //
	Username      interface{} //
	Password      interface{} //
	PasswordSalt  interface{} //
	LastLoginIp   interface{} //
	LastLoginTime interface{} //
	IsDelete      interface{} //
}
