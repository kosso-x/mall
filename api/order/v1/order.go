package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderSubmitReq struct {
	g.Meta       `path:"/submit" tags:"order" method:"post" summary:"提交订单"`
	AddressId    int
	Postscript   string
	FreightPrice int
	ActualPrice  float32
	OfflinePay   int
}

type OrderSubmitRes struct {
	g.Meta    `mime:"json"`
	OrderInfo *model.OrderInfo `json:"OrderInfo"`
}

type OrderListReq struct {
	g.Meta   `path:"/list" tags:"order" method:"get" summary:"订单列表"`
	ShowType int
	Size     int
	Page     int
}

type OrderListRes struct {
	g.Meta `mime:"json"`
	*model.OrderList
}

type OrderDetailReq struct {
	g.Meta  `path:"/detail" tags:"order" method:"get" summary:"订单详情"`
	OrderId int
}

type OrderDetailRes struct {
	g.Meta `mime:"json"`
	*model.OrderDetail
}

type OrderDeleteReq struct {
	g.Meta  `path:"/delete" tags:"order" method:"post" summary:"订单删除"`
	OrderId int
}

type OrderDeleteRes struct {
	g.Meta    `mime:"json"`
	Is_delete int `json:"is_delete"`
}

type OrderCancelReq struct {
	g.Meta  `path:"/cancel" tags:"order" method:"post" summary:"提交订单"`
	OrderId int
}

type OrderCancelRes struct {
	g.Meta   `mime:"json"`
	Canceled int `json:"canceled"`
}

type OrderConfirmReq struct {
	g.Meta  `path:"/confirm" tags:"order" method:"post" summary:"物流详情"`
	OrderId int
}

type OrderConfirmRes struct {
	g.Meta    `mime:"json"`
	Confirmed int `json:"confirmed"`
}

type OrderCountReq struct {
	g.Meta   `path:"/count" tags:"order" method:"get" summary:"获取订单数"`
	ShowType int
}

type OrderCountRes struct {
	g.Meta   `mime:"json"`
	AllCount int `json:"allCount"`
}

type OrderOrderCountReq struct {
	g.Meta `path:"/orderCount" tags:"order" method:"get" summary:"我的页面获取订单数状态"`
}

type OrderOrderCountRes struct {
	g.Meta `mime:"json"`
	*model.OrderCountInfo
}

type OrderExpressReq struct {
	g.Meta `path:"/express" tags:"order" method:"get" summary:"物流信息"`
}

type OrderExpressRes struct {
	g.Meta `mime:"json"`
}

type OrderOrderGoodsReq struct {
	g.Meta  `path:"/orderGoods" tags:"order" method:"get" summary:"获取checkout页面的商品列表"`
	OrderId int
}

type OrderOrderGoodsRes struct {
	g.Meta `mime:"json"`
	Res    interface{} `json:"data"`
}
