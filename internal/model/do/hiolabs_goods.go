// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsGoods is the golang structure of table hiolabs_goods for DAO operations like Where/Data.
type HiolabsGoods struct {
	g.Meta            `orm:"table:hiolabs_goods, do:true"`
	Id                interface{} //
	CategoryId        interface{} //
	IsOnSale          interface{} //
	Name              interface{} //
	GoodsNumber       interface{} //
	SellVolume        interface{} // 销售量
	Keywords          interface{} //
	RetailPrice       interface{} // 零售价格
	MinRetailPrice    interface{} //
	CostPrice         interface{} //
	MinCostPrice      interface{} //
	GoodsBrief        interface{} //
	GoodsDesc         interface{} //
	SortOrder         interface{} //
	IsIndex           interface{} //
	IsNew             interface{} //
	GoodsUnit         interface{} // 商品单位
	HttpsPicUrl       interface{} // 商品https图
	ListPicUrl        interface{} // 商品列表图
	FreightTemplateId interface{} //
	FreightType       interface{} //
	IsDelete          interface{} //
	HasGallery        interface{} //
	HasDone           interface{} //
}
