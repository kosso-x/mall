package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RegionListReq struct {
	g.Meta   `path:"/list" tags:"region" method:"get" summary:"获取区域列表"`
	ParentId int
}

type RegionListRes struct {
	g.Meta           `mime:"json"`
	model.RegionList `json:"regionList"`
}
