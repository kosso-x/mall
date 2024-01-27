// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsShipperDao is the data access object for table hiolabs_shipper.
type HiolabsShipperDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns HiolabsShipperColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsShipperColumns defines and stores column names for table hiolabs_shipper.
type HiolabsShipperColumns struct {
	Id           string //
	Name         string // 快递公司名称
	Code         string // 快递公司代码
	SortOrder    string // 排序
	MonthCode    string //
	CustomerName string //
	Enabled      string //
}

// hiolabsShipperColumns holds the columns for table hiolabs_shipper.
var hiolabsShipperColumns = HiolabsShipperColumns{
	Id:           "id",
	Name:         "name",
	Code:         "code",
	SortOrder:    "sort_order",
	MonthCode:    "MonthCode",
	CustomerName: "CustomerName",
	Enabled:      "enabled",
}

// NewHiolabsShipperDao creates and returns a new DAO object for table data access.
func NewHiolabsShipperDao() *HiolabsShipperDao {
	return &HiolabsShipperDao{
		group:   "default",
		table:   "hiolabs_shipper",
		columns: hiolabsShipperColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsShipperDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsShipperDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsShipperDao) Columns() HiolabsShipperColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsShipperDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsShipperDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsShipperDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
