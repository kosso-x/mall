package tableOperation

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var OrderStatusText = gmap.NewFrom(g.MapAnyAny{
	101: "未付款",
	102: "已取消",
	103: "已取消(系统)",
	201: "已付款",
	202: "订单取消，退款中",
	203: "已退款",
	301: "已发货",
	302: "已收货",
	303: "已收货(系统)",
	401: "已完成",
	801: "拼团中,未付款",
	802: "拼团中,已付款",
})

func OrderStatusEnum(status interface{}) (value string, found bool) {
	v, found := OrderStatusText.Search(gconv.Int(status))
	value = gconv.String(v)

	return
}
