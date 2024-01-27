// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsProduct is the golang structure of table hiolabs_product for DAO operations like Where/Data.
type HiolabsProduct struct {
	g.Meta                `orm:"table:hiolabs_product, do:true"`
	Id                    interface{} //
	GoodsId               interface{} //
	GoodsSpecificationIds interface{} //
	GoodsSn               interface{} //
	GoodsNumber           interface{} //
	RetailPrice           interface{} //
	Cost                  interface{} // 成本
	GoodsWeight           interface{} // 重量
	HasChange             interface{} //
	GoodsName             interface{} //
	IsOnSale              interface{} //
	IsDelete              interface{} //
}
