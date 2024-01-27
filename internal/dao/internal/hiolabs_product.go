// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsProductDao is the data access object for table hiolabs_product.
type HiolabsProductDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns HiolabsProductColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsProductColumns defines and stores column names for table hiolabs_product.
type HiolabsProductColumns struct {
	Id                    string //
	GoodsId               string //
	GoodsSpecificationIds string //
	GoodsSn               string //
	GoodsNumber           string //
	RetailPrice           string //
	Cost                  string // 成本
	GoodsWeight           string // 重量
	HasChange             string //
	GoodsName             string //
	IsOnSale              string //
	IsDelete              string //
}

// hiolabsProductColumns holds the columns for table hiolabs_product.
var hiolabsProductColumns = HiolabsProductColumns{
	Id:                    "id",
	GoodsId:               "goods_id",
	GoodsSpecificationIds: "goods_specification_ids",
	GoodsSn:               "goods_sn",
	GoodsNumber:           "goods_number",
	RetailPrice:           "retail_price",
	Cost:                  "cost",
	GoodsWeight:           "goods_weight",
	HasChange:             "has_change",
	GoodsName:             "goods_name",
	IsOnSale:              "is_on_sale",
	IsDelete:              "is_delete",
}

// NewHiolabsProductDao creates and returns a new DAO object for table data access.
func NewHiolabsProductDao() *HiolabsProductDao {
	return &HiolabsProductDao{
		group:   "default",
		table:   "hiolabs_product",
		columns: hiolabsProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsProductDao) Columns() HiolabsProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
