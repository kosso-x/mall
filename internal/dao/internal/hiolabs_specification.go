// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsSpecificationDao is the data access object for table hiolabs_specification.
type HiolabsSpecificationDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns HiolabsSpecificationColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsSpecificationColumns defines and stores column names for table hiolabs_specification.
type HiolabsSpecificationColumns struct {
	Id        string //
	Name      string //
	SortOrder string //
	Memo      string // 说明
}

// hiolabsSpecificationColumns holds the columns for table hiolabs_specification.
var hiolabsSpecificationColumns = HiolabsSpecificationColumns{
	Id:        "id",
	Name:      "name",
	SortOrder: "sort_order",
	Memo:      "memo",
}

// NewHiolabsSpecificationDao creates and returns a new DAO object for table data access.
func NewHiolabsSpecificationDao() *HiolabsSpecificationDao {
	return &HiolabsSpecificationDao{
		group:   "default",
		table:   "hiolabs_specification",
		columns: hiolabsSpecificationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsSpecificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsSpecificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsSpecificationDao) Columns() HiolabsSpecificationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsSpecificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsSpecificationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsSpecificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
