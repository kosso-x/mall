// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplateDetailDao is the data access object for table hiolabs_freight_template_detail.
type HiolabsFreightTemplateDetailDao struct {
	table   string                              // table is the underlying table name of the DAO.
	group   string                              // group is the database configuration group name of current DAO.
	columns HiolabsFreightTemplateDetailColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsFreightTemplateDetailColumns defines and stores column names for table hiolabs_freight_template_detail.
type HiolabsFreightTemplateDetailColumns struct {
	Id         string //
	TemplateId string //
	GroupId    string //
	Area       string // 0位默认，
	IsDelete   string //
}

// hiolabsFreightTemplateDetailColumns holds the columns for table hiolabs_freight_template_detail.
var hiolabsFreightTemplateDetailColumns = HiolabsFreightTemplateDetailColumns{
	Id:         "id",
	TemplateId: "template_id",
	GroupId:    "group_id",
	Area:       "area",
	IsDelete:   "is_delete",
}

// NewHiolabsFreightTemplateDetailDao creates and returns a new DAO object for table data access.
func NewHiolabsFreightTemplateDetailDao() *HiolabsFreightTemplateDetailDao {
	return &HiolabsFreightTemplateDetailDao{
		group:   "default",
		table:   "hiolabs_freight_template_detail",
		columns: hiolabsFreightTemplateDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsFreightTemplateDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsFreightTemplateDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsFreightTemplateDetailDao) Columns() HiolabsFreightTemplateDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsFreightTemplateDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsFreightTemplateDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsFreightTemplateDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
