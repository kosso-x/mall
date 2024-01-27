// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsGoodsSpecificationDao is the data access object for table hiolabs_goods_specification.
type HiolabsGoodsSpecificationDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns HiolabsGoodsSpecificationColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsGoodsSpecificationColumns defines and stores column names for table hiolabs_goods_specification.
type HiolabsGoodsSpecificationColumns struct {
	Id              string //
	GoodsId         string //
	SpecificationId string //
	Value           string //
	PicUrl          string //
	IsDelete        string //
}

// hiolabsGoodsSpecificationColumns holds the columns for table hiolabs_goods_specification.
var hiolabsGoodsSpecificationColumns = HiolabsGoodsSpecificationColumns{
	Id:              "id",
	GoodsId:         "goods_id",
	SpecificationId: "specification_id",
	Value:           "value",
	PicUrl:          "pic_url",
	IsDelete:        "is_delete",
}

// NewHiolabsGoodsSpecificationDao creates and returns a new DAO object for table data access.
func NewHiolabsGoodsSpecificationDao() *HiolabsGoodsSpecificationDao {
	return &HiolabsGoodsSpecificationDao{
		group:   "default",
		table:   "hiolabs_goods_specification",
		columns: hiolabsGoodsSpecificationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsGoodsSpecificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsGoodsSpecificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsGoodsSpecificationDao) Columns() HiolabsGoodsSpecificationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsGoodsSpecificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsGoodsSpecificationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsGoodsSpecificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
