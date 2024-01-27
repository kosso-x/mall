// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAddress is the golang structure of table hiolabs_address for DAO operations like Where/Data.
type HiolabsAddress struct {
	g.Meta     `orm:"table:hiolabs_address, do:true"`
	Id         interface{} //
	Name       interface{} //
	UserId     interface{} //
	CountryId  interface{} //
	ProvinceId interface{} //
	CityId     interface{} //
	DistrictId interface{} //
	Address    interface{} //
	Mobile     interface{} //
	IsDefault  interface{} //
	IsDelete   interface{} //
}
