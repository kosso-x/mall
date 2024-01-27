// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsGoodsGalleryDao is the data access object for table hiolabs_goods_gallery.
type HiolabsGoodsGalleryDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns HiolabsGoodsGalleryColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsGoodsGalleryColumns defines and stores column names for table hiolabs_goods_gallery.
type HiolabsGoodsGalleryColumns struct {
	Id        string //
	GoodsId   string //
	ImgUrl    string //
	ImgDesc   string //
	SortOrder string //
	IsDelete  string //
}

// hiolabsGoodsGalleryColumns holds the columns for table hiolabs_goods_gallery.
var hiolabsGoodsGalleryColumns = HiolabsGoodsGalleryColumns{
	Id:        "id",
	GoodsId:   "goods_id",
	ImgUrl:    "img_url",
	ImgDesc:   "img_desc",
	SortOrder: "sort_order",
	IsDelete:  "is_delete",
}

// NewHiolabsGoodsGalleryDao creates and returns a new DAO object for table data access.
func NewHiolabsGoodsGalleryDao() *HiolabsGoodsGalleryDao {
	return &HiolabsGoodsGalleryDao{
		group:   "default",
		table:   "hiolabs_goods_gallery",
		columns: hiolabsGoodsGalleryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsGoodsGalleryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsGoodsGalleryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsGoodsGalleryDao) Columns() HiolabsGoodsGalleryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsGoodsGalleryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsGoodsGalleryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsGoodsGalleryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
