// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplateGroupDao is the data access object for table hiolabs_freight_template_group.
type HiolabsFreightTemplateGroupDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of current DAO.
	columns HiolabsFreightTemplateGroupColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsFreightTemplateGroupColumns defines and stores column names for table hiolabs_freight_template_group.
type HiolabsFreightTemplateGroupColumns struct {
	Id           string //
	TemplateId   string //
	IsDefault    string // 默认，area:0
	Area         string // 0位默认，
	Start        string //
	StartFee     string //
	Add          string //
	AddFee       string //
	FreeByNumber string // 0没有设置
	FreeByMoney  string // 0没设置
	IsDelete     string //
}

// hiolabsFreightTemplateGroupColumns holds the columns for table hiolabs_freight_template_group.
var hiolabsFreightTemplateGroupColumns = HiolabsFreightTemplateGroupColumns{
	Id:           "id",
	TemplateId:   "template_id",
	IsDefault:    "is_default",
	Area:         "area",
	Start:        "start",
	StartFee:     "start_fee",
	Add:          "add",
	AddFee:       "add_fee",
	FreeByNumber: "free_by_number",
	FreeByMoney:  "free_by_money",
	IsDelete:     "is_delete",
}

// NewHiolabsFreightTemplateGroupDao creates and returns a new DAO object for table data access.
func NewHiolabsFreightTemplateGroupDao() *HiolabsFreightTemplateGroupDao {
	return &HiolabsFreightTemplateGroupDao{
		group:   "default",
		table:   "hiolabs_freight_template_group",
		columns: hiolabsFreightTemplateGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsFreightTemplateGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsFreightTemplateGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsFreightTemplateGroupDao) Columns() HiolabsFreightTemplateGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsFreightTemplateGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsFreightTemplateGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsFreightTemplateGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
