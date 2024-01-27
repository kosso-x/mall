// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsGoodsSpecification is the golang structure of table hiolabs_goods_specification for DAO operations like Where/Data.
type HiolabsGoodsSpecification struct {
	g.Meta          `orm:"table:hiolabs_goods_specification, do:true"`
	Id              interface{} //
	GoodsId         interface{} //
	SpecificationId interface{} //
	Value           interface{} //
	PicUrl          interface{} //
	IsDelete        interface{} //
}
