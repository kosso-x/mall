// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsOrderDao is the data access object for table hiolabs_order.
type HiolabsOrderDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns HiolabsOrderColumns // columns contains all the column names of Table for convenient usage.
}

// HiolabsOrderColumns defines and stores column names for table hiolabs_order.
type HiolabsOrderColumns struct {
	Id             string //
	OrderSn        string //
	UserId         string //
	OrderStatus    string // 101：未付款、102：已取消、103已取消(系统)、201：已付款、202：订单取消，退款中、203：已退款、301：已发货、302：已收货、303：已收货(系统)、401：已完成、801：拼团中,未付款、802：拼团中,已付款
	OfflinePay     string // 线下支付订单标志
	ShippingStatus string // 0未发货，1已发货
	PrintStatus    string //
	PayStatus      string //
	Consignee      string //
	Country        string //
	Province       string //
	City           string //
	District       string //
	Address        string //
	PrintInfo      string //
	Mobile         string //
	Postscript     string //
	AdminMemo      string //
	ShippingFee    string // 免邮的商品的邮费，这个在退款时不能退给客户
	PayName        string //
	PayId          string //
	ChangePrice    string // 0没改价，不等于0改过价格，这里记录原始的价格
	ActualPrice    string // 实际需要支付的金额
	OrderPrice     string // 订单总价
	GoodsPrice     string // 商品总价
	AddTime        string //
	PayTime        string // 付款时间
	ShippingTime   string // 发货时间
	ConfirmTime    string // 确认时间
	DealdoneTime   string // 成交时间，用户评论或自动好评时间
	FreightPrice   string // 配送费用
	ExpressValue   string // 顺丰保价金额
	Remark         string //
	OrderType      string // 订单类型：0普通，1秒杀，2团购，3返现订单,7充值，8会员
	IsDelete       string // 订单删除标志
}

// hiolabsOrderColumns holds the columns for table hiolabs_order.
var hiolabsOrderColumns = HiolabsOrderColumns{
	Id:             "id",
	OrderSn:        "order_sn",
	UserId:         "user_id",
	OrderStatus:    "order_status",
	OfflinePay:     "offline_pay",
	ShippingStatus: "shipping_status",
	PrintStatus:    "print_status",
	PayStatus:      "pay_status",
	Consignee:      "consignee",
	Country:        "country",
	Province:       "province",
	City:           "city",
	District:       "district",
	Address:        "address",
	PrintInfo:      "print_info",
	Mobile:         "mobile",
	Postscript:     "postscript",
	AdminMemo:      "admin_memo",
	ShippingFee:    "shipping_fee",
	PayName:        "pay_name",
	PayId:          "pay_id",
	ChangePrice:    "change_price",
	ActualPrice:    "actual_price",
	OrderPrice:     "order_price",
	GoodsPrice:     "goods_price",
	AddTime:        "add_time",
	PayTime:        "pay_time",
	ShippingTime:   "shipping_time",
	ConfirmTime:    "confirm_time",
	DealdoneTime:   "dealdone_time",
	FreightPrice:   "freight_price",
	ExpressValue:   "express_value",
	Remark:         "remark",
	OrderType:      "order_type",
	IsDelete:       "is_delete",
}

// NewHiolabsOrderDao creates and returns a new DAO object for table data access.
func NewHiolabsOrderDao() *HiolabsOrderDao {
	return &HiolabsOrderDao{
		group:   "default",
		table:   "hiolabs_order",
		columns: hiolabsOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HiolabsOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HiolabsOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HiolabsOrderDao) Columns() HiolabsOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HiolabsOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HiolabsOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HiolabsOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
