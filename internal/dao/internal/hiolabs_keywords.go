// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsKeywordsDao is the data access object for table hiolabs_keywords.
type HiolabsKeywordsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns HiolabsKeywordsColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsKeywordsColumns defines and stores column names for table hiolabs_keywords.
type HiolabsKeywordsColumns struct {
	Keyword   string //
	IsHot     string //
	IsDefault string //
	IsShow    string //
	SortOrder string //
	SchemeUrl string // 关键词的跳转链接
	Id        string //
	Type      string //
}

// hiolabsKeywordsColumns holds the columns for table hiolabs_keywords.
var hiolabsKeywordsColumns = HiolabsKeywordsColumns{
	Keyword:   "keyword",
	IsHot:     "is_hot",
	IsDefault: "is_default",
	IsShow:    "is_show",
	SortOrder: "sort_order",
	SchemeUrl: "scheme _url",
	Id:        "id",
	Type:      "type",
}

// NewHiolabsKeywordsDao creates and returns a new DAO object for table data access.
func NewHiolabsKeywordsDao() *HiolabsKeywordsDao {
	return &HiolabsKeywordsDao{
		group:   "default",
		table:   "hiolabs_keywords",
		columns: hiolabsKeywordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsKeywordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsKeywordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsKeywordsDao) Columns() HiolabsKeywordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsKeywordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsKeywordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsKeywordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
