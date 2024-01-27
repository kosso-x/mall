// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsSettingsDao is the data access object for table hiolabs_settings.
type HiolabsSettingsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns HiolabsSettingsColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsSettingsColumns defines and stores column names for table hiolabs_settings.
type HiolabsSettingsColumns struct {
	Id                 string //
	AutoDelivery       string //
	Name               string //
	Tel                string //
	ProvinceName       string //
	CityName           string //
	ExpAreaName        string //
	Address            string //
	DiscoveryImgHeight string //
	DiscoveryImg       string //
	GoodsId            string //
	CityId             string //
	ProvinceId         string //
	DistrictId         string //
	Countdown          string // 10分钟倒计时
	Reset              string //
}

// hiolabsSettingsColumns holds the columns for table hiolabs_settings.
var hiolabsSettingsColumns = HiolabsSettingsColumns{
	Id:                 "id",
	AutoDelivery:       "autoDelivery",
	Name:               "Name",
	Tel:                "Tel",
	ProvinceName:       "ProvinceName",
	CityName:           "CityName",
	ExpAreaName:        "ExpAreaName",
	Address:            "Address",
	DiscoveryImgHeight: "discovery_img_height",
	DiscoveryImg:       "discovery_img",
	GoodsId:            "goods_id",
	CityId:             "city_id",
	ProvinceId:         "province_id",
	DistrictId:         "district_id",
	Countdown:          "countdown",
	Reset:              "reset",
}

// NewHiolabsSettingsDao creates and returns a new DAO object for table data access.
func NewHiolabsSettingsDao() *HiolabsSettingsDao {
	return &HiolabsSettingsDao{
		group:   "default",
		table:   "hiolabs_settings",
		columns: hiolabsSettingsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsSettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsSettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsSettingsDao) Columns() HiolabsSettingsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsSettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsSettingsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsSettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
