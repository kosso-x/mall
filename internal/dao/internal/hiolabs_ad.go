// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAdDao is the data access object for table hiolabs_ad.
type HiolabsAdDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns HiolabsAdColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsAdColumns defines and stores column names for table hiolabs_ad.
type HiolabsAdColumns struct {
	Id        string //
	LinkType  string // 0商品，1链接
	Link      string //
	GoodsId   string //
	ImageUrl  string //
	EndTime   string //
	Enabled   string //
	SortOrder string //
	IsDelete  string //
}

// hiolabsAdColumns holds the columns for table hiolabs_ad.
var hiolabsAdColumns = HiolabsAdColumns{
	Id:        "id",
	LinkType:  "link_type",
	Link:      "link",
	GoodsId:   "goods_id",
	ImageUrl:  "image_url",
	EndTime:   "end_time",
	Enabled:   "enabled",
	SortOrder: "sort_order",
	IsDelete:  "is_delete",
}

// NewHiolabsAdDao creates and returns a new DAO object for table data access.
func NewHiolabsAdDao() *HiolabsAdDao {
	return &HiolabsAdDao{
		group:   "default",
		table:   "hiolabs_ad",
		columns: hiolabsAdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsAdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsAdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsAdDao) Columns() HiolabsAdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsAdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsAdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsAdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
