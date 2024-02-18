package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type FootprintListReq struct {
	g.Meta `path:"/list" tags:"footprint" method:"get" summary:"足迹列表"`
	Page   int
	Size   int
}

type FootprintListRes struct {
	g.Meta `mime:"json"`
	*model.FootprintList
}

type FootprintDeleteReq struct {
	g.Meta      `path:"/delete" tags:"footprint" method:"post" summary:"删除足迹"`
	FootprintId int
}

type FootprintDeleteRes struct {
	g.Meta   `mime:"json"`
	IsDelete int `json:"is_delete"`
}
