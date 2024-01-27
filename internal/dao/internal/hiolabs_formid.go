// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFormidDao is the data access object for table hiolabs_formid.
type HiolabsFormidDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns HiolabsFormidColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsFormidColumns defines and stores column names for table hiolabs_formid.
type HiolabsFormidColumns struct {
	Id       string //
	UserId   string //
	OrderId  string //
	FormId   string //
	AddTime  string // 添加时间
	UseTimes string // 使用次数
}

// hiolabsFormidColumns holds the columns for table hiolabs_formid.
var hiolabsFormidColumns = HiolabsFormidColumns{
	Id:       "id",
	UserId:   "user_id",
	OrderId:  "order_id",
	FormId:   "form_id",
	AddTime:  "add_time",
	UseTimes: "use_times",
}

// NewHiolabsFormidDao creates and returns a new DAO object for table data access.
func NewHiolabsFormidDao() *HiolabsFormidDao {
	return &HiolabsFormidDao{
		group:   "default",
		table:   "hiolabs_formid",
		columns: hiolabsFormidColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsFormidDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsFormidDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsFormidDao) Columns() HiolabsFormidColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsFormidDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsFormidDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsFormidDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
