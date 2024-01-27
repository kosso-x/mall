// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsFreightTemplateDao is the data access object for table hiolabs_freight_template.
type HiolabsFreightTemplateDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns HiolabsFreightTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsFreightTemplateColumns defines and stores column names for table hiolabs_freight_template.
type HiolabsFreightTemplateColumns struct {
	Id           string //
	Name         string // 运费模板名称
	PackagePrice string // 包装费用
	FreightType  string // 0按件，1按重量
	IsDelete     string //
}

// hiolabsFreightTemplateColumns holds the columns for table hiolabs_freight_template.
var hiolabsFreightTemplateColumns = HiolabsFreightTemplateColumns{
	Id:           "id",
	Name:         "name",
	PackagePrice: "package_price",
	FreightType:  "freight_type",
	IsDelete:     "is_delete",
}

// NewHiolabsFreightTemplateDao creates and returns a new DAO object for table data access.
func NewHiolabsFreightTemplateDao() *HiolabsFreightTemplateDao {
	return &HiolabsFreightTemplateDao{
		group:   "default",
		table:   "hiolabs_freight_template",
		columns: hiolabsFreightTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsFreightTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsFreightTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsFreightTemplateDao) Columns() HiolabsFreightTemplateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsFreightTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsFreightTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsFreightTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
