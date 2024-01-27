// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrderGoodsDao is the data access object for table hiolabs_order_goods.
type HiolabsOrderGoodsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns HiolabsOrderGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsOrderGoodsColumns defines and stores column names for table hiolabs_order_goods.
type HiolabsOrderGoodsColumns struct {
	Id                        string //
	OrderId                   string //
	GoodsId                   string //
	GoodsName                 string //
	GoodsAka                  string //
	ProductId                 string //
	Number                    string //
	RetailPrice               string //
	GoodsSpecifitionNameValue string //
	GoodsSpecifitionIds       string //
	ListPicUrl                string //
	UserId                    string //
	IsDelete                  string // 是否删除标志
}

// hiolabsOrderGoodsColumns holds the columns for table hiolabs_order_goods.
var hiolabsOrderGoodsColumns = HiolabsOrderGoodsColumns{
	Id:                        "id",
	OrderId:                   "order_id",
	GoodsId:                   "goods_id",
	GoodsName:                 "goods_name",
	GoodsAka:                  "goods_aka",
	ProductId:                 "product_id",
	Number:                    "number",
	RetailPrice:               "retail_price",
	GoodsSpecifitionNameValue: "goods_specifition_name_value",
	GoodsSpecifitionIds:       "goods_specifition_ids",
	ListPicUrl:                "list_pic_url",
	UserId:                    "user_id",
	IsDelete:                  "is_delete",
}

// NewHiolabsOrderGoodsDao creates and returns a new DAO object for table data access.
func NewHiolabsOrderGoodsDao() *HiolabsOrderGoodsDao {
	return &HiolabsOrderGoodsDao{
		group:   "default",
		table:   "hiolabs_order_goods",
		columns: hiolabsOrderGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsOrderGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsOrderGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsOrderGoodsDao) Columns() HiolabsOrderGoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsOrderGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsOrderGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsOrderGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
