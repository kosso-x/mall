package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 添加到购物车
type CartAddReq struct {
	g.Meta    `path:"/add" tags:"cart" method:"post" summary:"添加到购物车"`
	AddType   int    `v:"required"`
	GoodsId   string `v:"required"`
	Number    int    `v:"required"`
	ProductId int    `v:"required"`
}
type CartAddRes struct {
	g.Meta `mime:"json"`
	*model.CartAddRes
}

// 获取购物车的数据
type CartIndexReq struct {
	g.Meta `path:"/index" tags:"cart" method:"get" summary:"获取购物车的数据"`
}
type CartIndexRes struct {
	g.Meta `mime:"json"`
	*model.CartAddRes
}

// 更新购物车的商品
type CartUpdateReq struct {
	g.Meta    `path:"/update" tags:"cart" method:"post" summary:"更新购物车的商品"`
	Id        int
	Number    int
	ProductId int
}
type CartUpdateRes struct {
	g.Meta `mime:"json"`
	*model.CartAddRes
}

// 删除购物车的商品
type CartDeleteReq struct {
	g.Meta     `path:"/delete" tags:"cart" method:"post" summary:"删除购物车的商品"`
	ProductIds []int
}
type CartDeleteRes struct {
	g.Meta `mime:"json"`
	*model.CartAddRes
}

// 选择或取消选择商品
type CartCheckReq struct {
	g.Meta     `path:"/checked" tags:"cart" method:"post" summary:"选择或取消选择商品"`
	ProductIds []int
	IsChecked  int
}
type CartCheckRes struct {
	g.Meta `mime:"json"`
	*model.CartAddRes
}

// 获取购物车商品件数
type CartGoodsCountReq struct {
	g.Meta `path:"/goodsCount" tags:"cart" method:"get" summary:"获取购物车商品件数"`
}
type CartGoodsCountRes struct {
	g.Meta `mime:"json"`
	*model.CartGoodsCount
}

// 下单前信息确认
type CartCheckoutReq struct {
	g.Meta    `path:"/checkout" tags:"cart" method:"get" summary:"下单前信息确认"`
	AddressId int
	AddType   int
	OrderFrom string
	Type      int
}
type CartCheckoutRes struct {
	g.Meta `mime:"json"`
	*model.CartCheckout
}
