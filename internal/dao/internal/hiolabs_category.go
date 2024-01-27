// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsCategoryDao is the data access object for table hiolabs_category.
type HiolabsCategoryDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns HiolabsCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsCategoryColumns defines and stores column names for table hiolabs_category.
type HiolabsCategoryColumns struct {
	Id         string //
	Name       string //
	Keywords   string //
	FrontDesc  string //
	ParentId   string //
	SortOrder  string //
	ShowIndex  string //
	IsShow     string //
	IconUrl    string //
	ImgUrl     string //
	Level      string //
	FrontName  string //
	PHeight    string //
	IsCategory string //
	IsChannel  string //
}

// hiolabsCategoryColumns holds the columns for table hiolabs_category.
var hiolabsCategoryColumns = HiolabsCategoryColumns{
	Id:         "id",
	Name:       "name",
	Keywords:   "keywords",
	FrontDesc:  "front_desc",
	ParentId:   "parent_id",
	SortOrder:  "sort_order",
	ShowIndex:  "show_index",
	IsShow:     "is_show",
	IconUrl:    "icon_url",
	ImgUrl:     "img_url",
	Level:      "level",
	FrontName:  "front_name",
	PHeight:    "p_height",
	IsCategory: "is_category",
	IsChannel:  "is_channel",
}

// NewHiolabsCategoryDao creates and returns a new DAO object for table data access.
func NewHiolabsCategoryDao() *HiolabsCategoryDao {
	return &HiolabsCategoryDao{
		group:   "default",
		table:   "hiolabs_category",
		columns: hiolabsCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsCategoryDao) Columns() HiolabsCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
