// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsCartDao is the data access object for table hiolabs_cart.
type HiolabsCartDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns HiolabsCartColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsCartColumns defines and stores column names for table hiolabs_cart.
type HiolabsCartColumns struct {
	Id                        string //
	UserId                    string //
	GoodsId                   string //
	GoodsSn                   string //
	ProductId                 string //
	GoodsName                 string //
	GoodsAka                  string //
	GoodsWeight               string // 重量
	AddPrice                  string // 加入购物车时的价格
	RetailPrice               string //
	Number                    string //
	GoodsSpecifitionNameValue string // 规格属性组成的字符串，用来显示用
	GoodsSpecifitionIds       string // product表对应的goods_specifition_ids
	Checked                   string //
	ListPicUrl                string //
	FreightTemplateId         string // 运费模板
	IsOnSale                  string // 0
	AddTime                   string //
	IsFast                    string // 1
	IsDelete                  string //
}

// hiolabsCartColumns holds the columns for table hiolabs_cart.
var hiolabsCartColumns = HiolabsCartColumns{
	Id:                        "id",
	UserId:                    "user_id",
	GoodsId:                   "goods_id",
	GoodsSn:                   "goods_sn",
	ProductId:                 "product_id",
	GoodsName:                 "goods_name",
	GoodsAka:                  "goods_aka",
	GoodsWeight:               "goods_weight",
	AddPrice:                  "add_price",
	RetailPrice:               "retail_price",
	Number:                    "number",
	GoodsSpecifitionNameValue: "goods_specifition_name_value",
	GoodsSpecifitionIds:       "goods_specifition_ids",
	Checked:                   "checked",
	ListPicUrl:                "list_pic_url",
	FreightTemplateId:         "freight_template_id",
	IsOnSale:                  "is_on_sale",
	AddTime:                   "add_time",
	IsFast:                    "is_fast",
	IsDelete:                  "is_delete",
}

// NewHiolabsCartDao creates and returns a new DAO object for table data access.
func NewHiolabsCartDao() *HiolabsCartDao {
	return &HiolabsCartDao{
		group:   "default",
		table:   "hiolabs_cart",
		columns: hiolabsCartColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsCartDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsCartDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsCartDao) Columns() HiolabsCartColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsCartDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsCartDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsCartDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
