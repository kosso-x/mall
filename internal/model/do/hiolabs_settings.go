// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsSettings is the golang structure of table hiolabs_settings for DAO operations like Where/Data.
type HiolabsSettings struct {
	g.Meta             `orm:"table:hiolabs_settings, do:true"`
	Id                 interface{} //
	AutoDelivery       interface{} //
	Name               interface{} //
	Tel                interface{} //
	ProvinceName       interface{} //
	CityName           interface{} //
	ExpAreaName        interface{} //
	Address            interface{} //
	DiscoveryImgHeight interface{} //
	DiscoveryImg       interface{} //
	GoodsId            interface{} //
	CityId             interface{} //
	ProvinceId         interface{} //
	DistrictId         interface{} //
	Countdown          interface{} // 10分钟倒计时
	Reset              interface{} //
}
