// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsAddressDao is the data access object for table hiolabs_address.
type HiolabsAddressDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns HiolabsAddressColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsAddressColumns defines and stores column names for table hiolabs_address.
type HiolabsAddressColumns struct {
	Id         string //
	Name       string //
	UserId     string //
	CountryId  string //
	ProvinceId string //
	CityId     string //
	DistrictId string //
	Address    string //
	Mobile     string //
	IsDefault  string //
	IsDelete   string //
}

// hiolabsAddressColumns holds the columns for table hiolabs_address.
var hiolabsAddressColumns = HiolabsAddressColumns{
	Id:         "id",
	Name:       "name",
	UserId:     "user_id",
	CountryId:  "country_id",
	ProvinceId: "province_id",
	CityId:     "city_id",
	DistrictId: "district_id",
	Address:    "address",
	Mobile:     "mobile",
	IsDefault:  "is_default",
	IsDelete:   "is_delete",
}

// NewHiolabsAddressDao creates and returns a new DAO object for table data access.
func NewHiolabsAddressDao() *HiolabsAddressDao {
	return &HiolabsAddressDao{
		group:   "default",
		table:   "hiolabs_address",
		columns: hiolabsAddressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsAddressDao) Columns() HiolabsAddressColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsAddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
