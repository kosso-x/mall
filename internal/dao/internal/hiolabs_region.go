// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsRegionDao is the data access object for table hiolabs_region.
type HiolabsRegionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns HiolabsRegionColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsRegionColumns defines and stores column names for table hiolabs_region.
type HiolabsRegionColumns struct {
	Id       string //
	ParentId string //
	Name     string //
	Type     string //
	AgencyId string //
	Area     string // 方位，根据这个定运费
	AreaCode string // 方位代码
	FarArea  string // 偏远地区
}

// hiolabsRegionColumns holds the columns for table hiolabs_region.
var hiolabsRegionColumns = HiolabsRegionColumns{
	Id:       "id",
	ParentId: "parent_id",
	Name:     "name",
	Type:     "type",
	AgencyId: "agency_id",
	Area:     "area",
	AreaCode: "area_code",
	FarArea:  "far_area",
}

// NewHiolabsRegionDao creates and returns a new DAO object for table data access.
func NewHiolabsRegionDao() *HiolabsRegionDao {
	return &HiolabsRegionDao{
		group:   "default",
		table:   "hiolabs_region",
		columns: hiolabsRegionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsRegionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsRegionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsRegionDao) Columns() HiolabsRegionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsRegionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsRegionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsRegionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
