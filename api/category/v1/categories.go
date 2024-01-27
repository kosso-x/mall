package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CategoriesReq struct {
	g.Meta `path:"/index" tags:"catalog" method:"get" summary:"获取全部商品分类"`
}
type CategoriesRes struct {
	g.Meta              `mime:"json"`
	*model.CategoryList `json:"categoryList"`
}

type CategoryCurrentReq struct {
	g.Meta `path:"/current" tags:"catalog" method:"get" summary:"获取当前选中商品分类"`
	Id     int `v:"required"`
}
type CategoryCurrentRes struct {
	g.Meta `mime:"json"`
	*model.CurrentCategory
}

type CategoryCurrentListReq struct {
	g.Meta `path:"/currentlist" tags:"catalog" method:"post" summary:"获取当前选中商品分类及相关商品"`
	Id     int
	Size   int
	Page   int
}

type CategoryCurrentListRes struct {
	g.Meta `mime:"json"`
	*model.CategoryCurrentList
}
