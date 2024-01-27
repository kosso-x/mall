// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsCart is the golang structure of table hiolabs_cart for DAO operations like Where/Data.
type HiolabsCart struct {
	g.Meta                    `orm:"table:hiolabs_cart, do:true"`
	Id                        interface{} //
	UserId                    interface{} //
	GoodsId                   interface{} //
	GoodsSn                   interface{} //
	ProductId                 interface{} //
	GoodsName                 interface{} //
	GoodsAka                  interface{} //
	GoodsWeight               interface{} // 重量
	AddPrice                  interface{} // 加入购物车时的价格
	RetailPrice               interface{} //
	Number                    interface{} //
	GoodsSpecifitionNameValue interface{} // 规格属性组成的字符串，用来显示用
	GoodsSpecifitionIds       interface{} // product表对应的goods_specifition_ids
	Checked                   interface{} //
	ListPicUrl                interface{} //
	FreightTemplateId         interface{} // 运费模板
	IsOnSale                  interface{} // 0
	AddTime                   interface{} //
	IsFast                    interface{} // 1
	IsDelete                  interface{} //
}
