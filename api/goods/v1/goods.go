package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 统计商品总数
type GoodsCountReq struct {
	g.Meta `path:"/count" tags:"goods" method:"get" summary:"统计商品总数"`
}

type GoodsCountRes struct {
	g.Meta     `mime:"json"`
	GoodsCount int `json:"goodsCount"`
}

// 获得商品的详情
type GoodsDetailReq struct {
	g.Meta `path:"/detail" tags:"goods" method:"get" summary:"获得商品的详情"`
	Id     int
}

type GoodsDetailRes struct {
	g.Meta `mime:"json"`
	*model.GoodsDetail
}

// 获得商品列表
type GoodsListReq struct {
	g.Meta `path:"/list" tags:"goods" method:"get" summary:"获得商品列表"`
}

type GoodsListRes struct {
	g.Meta               `mime:"json"`
	*model.GoodsListInfo `json:"goodsListInfo"`
}

// 获得商品的详情
type GoodsShareReq struct {
	g.Meta `path:"/goodsShare" tags:"goods" method:"get" summary:"获得商品的详情"`
	Id     int
}

type GoodsShareRes struct {
	g.Meta `mime:"json"`
	*model.GoodsShare
}
