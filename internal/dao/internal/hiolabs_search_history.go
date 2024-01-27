// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsSearchHistoryDao is the data access object for table hiolabs_search_history.
type HiolabsSearchHistoryDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns HiolabsSearchHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsSearchHistoryColumns defines and stores column names for table hiolabs_search_history.
type HiolabsSearchHistoryColumns struct {
	Id      string //
	Keyword string //
	From    string // 搜索来源，如PC、小程序、APP等
	AddTime string // 搜索时间
	UserId  string //
}

// hiolabsSearchHistoryColumns holds the columns for table hiolabs_search_history.
var hiolabsSearchHistoryColumns = HiolabsSearchHistoryColumns{
	Id:      "id",
	Keyword: "keyword",
	From:    "from",
	AddTime: "add_time",
	UserId:  "user_id",
}

// NewHiolabsSearchHistoryDao creates and returns a new DAO object for table data access.
func NewHiolabsSearchHistoryDao() *HiolabsSearchHistoryDao {
	return &HiolabsSearchHistoryDao{
		group:   "default",
		table:   "hiolabs_search_history",
		columns: hiolabsSearchHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsSearchHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsSearchHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsSearchHistoryDao) Columns() HiolabsSearchHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsSearchHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsSearchHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsSearchHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
