// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsNoticeDao is the data access object for table hiolabs_notice.
type HiolabsNoticeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns HiolabsNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsNoticeColumns defines and stores column names for table hiolabs_notice.
type HiolabsNoticeColumns struct {
	Id       string //
	Content  string //
	EndTime  string //
	IsDelete string //
}

// hiolabsNoticeColumns holds the columns for table hiolabs_notice.
var hiolabsNoticeColumns = HiolabsNoticeColumns{
	Id:       "id",
	Content:  "content",
	EndTime:  "end_time",
	IsDelete: "is_delete",
}

// NewHiolabsNoticeDao creates and returns a new DAO object for table data access.
func NewHiolabsNoticeDao() *HiolabsNoticeDao {
	return &HiolabsNoticeDao{
		group:   "default",
		table:   "hiolabs_notice",
		columns: hiolabsNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsNoticeDao) Columns() HiolabsNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
