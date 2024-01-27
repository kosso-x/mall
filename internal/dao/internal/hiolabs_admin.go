// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAdminDao is the data access object for table hiolabs_admin.
type HiolabsAdminDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns HiolabsAdminColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsAdminColumns defines and stores column names for table hiolabs_admin.
type HiolabsAdminColumns struct {
	Id            string //
	Username      string //
	Password      string //
	PasswordSalt  string //
	LastLoginIp   string //
	LastLoginTime string //
	IsDelete      string //
}

// hiolabsAdminColumns holds the columns for table hiolabs_admin.
var hiolabsAdminColumns = HiolabsAdminColumns{
	Id:            "id",
	Username:      "username",
	Password:      "password",
	PasswordSalt:  "password_salt",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	IsDelete:      "is_delete",
}

// NewHiolabsAdminDao creates and returns a new DAO object for table data access.
func NewHiolabsAdminDao() *HiolabsAdminDao {
	return &HiolabsAdminDao{
		group:   "default",
		table:   "hiolabs_admin",
		columns: hiolabsAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsAdminDao) Columns() HiolabsAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
