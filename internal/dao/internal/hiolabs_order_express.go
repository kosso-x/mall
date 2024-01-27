// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrderExpressDao is the data access object for table hiolabs_order_express.
type HiolabsOrderExpressDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns HiolabsOrderExpressColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsOrderExpressColumns defines and stores column names for table hiolabs_order_express.
type HiolabsOrderExpressColumns struct {
	Id           string //
	OrderId      string //
	ShipperId    string //
	ShipperName  string // 物流公司名称
	ShipperCode  string // 物流公司代码
	LogisticCode string // 快递单号
	Traces       string // 物流跟踪信息
	IsFinish     string //
	RequestCount string // 总查询次数
	RequestTime  string // 最近一次向第三方查询物流信息时间
	AddTime      string // 添加时间
	UpdateTime   string // 更新时间
	ExpressType  string //
	RegionCode   string // 快递的地区编码，如杭州571
}

// hiolabsOrderExpressColumns holds the columns for table hiolabs_order_express.
var hiolabsOrderExpressColumns = HiolabsOrderExpressColumns{
	Id:           "id",
	OrderId:      "order_id",
	ShipperId:    "shipper_id",
	ShipperName:  "shipper_name",
	ShipperCode:  "shipper_code",
	LogisticCode: "logistic_code",
	Traces:       "traces",
	IsFinish:     "is_finish",
	RequestCount: "request_count",
	RequestTime:  "request_time",
	AddTime:      "add_time",
	UpdateTime:   "update_time",
	ExpressType:  "express_type",
	RegionCode:   "region_code",
}

// NewHiolabsOrderExpressDao creates and returns a new DAO object for table data access.
func NewHiolabsOrderExpressDao() *HiolabsOrderExpressDao {
	return &HiolabsOrderExpressDao{
		group:   "default",
		table:   "hiolabs_order_express",
		columns: hiolabsOrderExpressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsOrderExpressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsOrderExpressDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsOrderExpressDao) Columns() HiolabsOrderExpressColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsOrderExpressDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsOrderExpressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsOrderExpressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
