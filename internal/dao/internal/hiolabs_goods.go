// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsGoodsDao is the data access object for table hiolabs_goods.
type HiolabsGoodsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns HiolabsGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsGoodsColumns defines and stores column names for table hiolabs_goods.
type HiolabsGoodsColumns struct {
	Id                string //
	CategoryId        string //
	IsOnSale          string //
	Name              string //
	GoodsNumber       string //
	SellVolume        string // 销售量
	Keywords          string //
	RetailPrice       string // 零售价格
	MinRetailPrice    string //
	CostPrice         string //
	MinCostPrice      string //
	GoodsBrief        string //
	GoodsDesc         string //
	SortOrder         string //
	IsIndex           string //
	IsNew             string //
	GoodsUnit         string // 商品单位
	HttpsPicUrl       string // 商品https图
	ListPicUrl        string // 商品列表图
	FreightTemplateId string //
	FreightType       string //
	IsDelete          string //
	HasGallery        string //
	HasDone           string //
}

// hiolabsGoodsColumns holds the columns for table hiolabs_goods.
var hiolabsGoodsColumns = HiolabsGoodsColumns{
	Id:                "id",
	CategoryId:        "category_id",
	IsOnSale:          "is_on_sale",
	Name:              "name",
	GoodsNumber:       "goods_number",
	SellVolume:        "sell_volume",
	Keywords:          "keywords",
	RetailPrice:       "retail_price",
	MinRetailPrice:    "min_retail_price",
	CostPrice:         "cost_price",
	MinCostPrice:      "min_cost_price",
	GoodsBrief:        "goods_brief",
	GoodsDesc:         "goods_desc",
	SortOrder:         "sort_order",
	IsIndex:           "is_index",
	IsNew:             "is_new",
	GoodsUnit:         "goods_unit",
	HttpsPicUrl:       "https_pic_url",
	ListPicUrl:        "list_pic_url",
	FreightTemplateId: "freight_template_id",
	FreightType:       "freight_type",
	IsDelete:          "is_delete",
	HasGallery:        "has_gallery",
	HasDone:           "has_done",
}

// NewHiolabsGoodsDao creates and returns a new DAO object for table data access.
func NewHiolabsGoodsDao() *HiolabsGoodsDao {
	return &HiolabsGoodsDao{
		group:   "default",
		table:   "hiolabs_goods",
		columns: hiolabsGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsGoodsDao) Columns() HiolabsGoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
