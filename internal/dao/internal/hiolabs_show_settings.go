// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsShowSettingsDao is the data access object for table hiolabs_show_settings.
type HiolabsShowSettingsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns HiolabsShowSettingsColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsShowSettingsColumns defines and stores column names for table hiolabs_show_settings.
type HiolabsShowSettingsColumns struct {
	Id             string //
	Banner         string // 滚动banner
	Channel        string // 是否开启menu,几个图标
	IndexBannerImg string // 首页的img图片是否显示
	Notice         string //
}

// hiolabsShowSettingsColumns holds the columns for table hiolabs_show_settings.
var hiolabsShowSettingsColumns = HiolabsShowSettingsColumns{
	Id:             "id",
	Banner:         "banner",
	Channel:        "channel",
	IndexBannerImg: "index_banner_img",
	Notice:         "notice",
}

// NewHiolabsShowSettingsDao creates and returns a new DAO object for table data access.
func NewHiolabsShowSettingsDao() *HiolabsShowSettingsDao {
	return &HiolabsShowSettingsDao{
		group:   "default",
		table:   "hiolabs_show_settings",
		columns: hiolabsShowSettingsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsShowSettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsShowSettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsShowSettingsDao) Columns() HiolabsShowSettingsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsShowSettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsShowSettingsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsShowSettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
