// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrderGoods is the golang structure of table hiolabs_order_goods for DAO operations like Where/Data.
type HiolabsOrderGoods struct {
	g.Meta                    `orm:"table:hiolabs_order_goods, do:true"`
	Id                        interface{} //
	OrderId                   interface{} //
	GoodsId                   interface{} //
	GoodsName                 interface{} //
	GoodsAka                  interface{} //
	ProductId                 interface{} //
	Number                    interface{} //
	RetailPrice               interface{} //
	GoodsSpecifitionNameValue interface{} //
	GoodsSpecifitionIds       interface{} //
	ListPicUrl                interface{} //
	UserId                    interface{} //
	IsDelete                  interface{} // 是否删除标志
}
