// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsUserDao is the data access object for table hiolabs_user.
type HiolabsUserDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns HiolabsUserColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsUserColumns defines and stores column names for table hiolabs_user.
type HiolabsUserColumns struct {
	Id            string //
	Nickname      string //
	Name          string //
	Username      string //
	Password      string //
	Gender        string //
	Birthday      string //
	RegisterTime  string //
	LastLoginTime string //
	LastLoginIp   string //
	Mobile        string //
	RegisterIp    string //
	Avatar        string //
	WeixinOpenid  string //
	NameMobile    string //
	Country       string //
	Province      string //
	City          string //
}

// hiolabsUserColumns holds the columns for table hiolabs_user.
var hiolabsUserColumns = HiolabsUserColumns{
	Id:            "id",
	Nickname:      "nickname",
	Name:          "name",
	Username:      "username",
	Password:      "password",
	Gender:        "gender",
	Birthday:      "birthday",
	RegisterTime:  "register_time",
	LastLoginTime: "last_login_time",
	LastLoginIp:   "last_login_ip",
	Mobile:        "mobile",
	RegisterIp:    "register_ip",
	Avatar:        "avatar",
	WeixinOpenid:  "weixin_openid",
	NameMobile:    "name_mobile",
	Country:       "country",
	Province:      "province",
	City:          "city",
}

// NewHiolabsUserDao creates and returns a new DAO object for table data access.
func NewHiolabsUserDao() *HiolabsUserDao {
	return &HiolabsUserDao{
		group:   "default",
		table:   "hiolabs_user",
		columns: hiolabsUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsUserDao) Columns() HiolabsUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
