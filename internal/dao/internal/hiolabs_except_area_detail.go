// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsExceptAreaDetailDao is the data access object for table hiolabs_except_area_detail.
type HiolabsExceptAreaDetailDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns HiolabsExceptAreaDetailColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsExceptAreaDetailColumns defines and stores column names for table hiolabs_except_area_detail.
type HiolabsExceptAreaDetailColumns struct {
	Id           string //
	ExceptAreaId string //
	Area         string // 0位默认，
	IsDelete     string //
}

// hiolabsExceptAreaDetailColumns holds the columns for table hiolabs_except_area_detail.
var hiolabsExceptAreaDetailColumns = HiolabsExceptAreaDetailColumns{
	Id:           "id",
	ExceptAreaId: "except_area_id",
	Area:         "area",
	IsDelete:     "is_delete",
}

// NewHiolabsExceptAreaDetailDao creates and returns a new DAO object for table data access.
func NewHiolabsExceptAreaDetailDao() *HiolabsExceptAreaDetailDao {
	return &HiolabsExceptAreaDetailDao{
		group:   "default",
		table:   "hiolabs_except_area_detail",
		columns: hiolabsExceptAreaDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsExceptAreaDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsExceptAreaDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsExceptAreaDetailDao) Columns() HiolabsExceptAreaDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsExceptAreaDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsExceptAreaDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsExceptAreaDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
