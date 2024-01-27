// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsExceptAreaDao is the data access object for table hiolabs_except_area.
type HiolabsExceptAreaDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns HiolabsExceptAreaColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsExceptAreaColumns defines and stores column names for table hiolabs_except_area.
type HiolabsExceptAreaColumns struct {
	Id       string //
	Content  string // 名称
	Area     string // 0位默认，
	IsDelete string //
}

// hiolabsExceptAreaColumns holds the columns for table hiolabs_except_area.
var hiolabsExceptAreaColumns = HiolabsExceptAreaColumns{
	Id:       "id",
	Content:  "content",
	Area:     "area",
	IsDelete: "is_delete",
}

// NewHiolabsExceptAreaDao creates and returns a new DAO object for table data access.
func NewHiolabsExceptAreaDao() *HiolabsExceptAreaDao {
	return &HiolabsExceptAreaDao{
		group:   "default",
		table:   "hiolabs_except_area",
		columns: hiolabsExceptAreaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsExceptAreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsExceptAreaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsExceptAreaDao) Columns() HiolabsExceptAreaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsExceptAreaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsExceptAreaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsExceptAreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
