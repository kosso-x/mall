// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAd is the golang structure of table hiolabs_ad for DAO operations like Where/Data.
type HiolabsAd struct {
	g.Meta    `orm:"table:hiolabs_ad, do:true"`
	Id        interface{} //
	LinkType  interface{} // 0商品，1链接
	Link      interface{} //
	GoodsId   interface{} //
	ImageUrl  interface{} //
	EndTime   interface{} //
	Enabled   interface{} //
	SortOrder interface{} //
	IsDelete  interface{} //
}
