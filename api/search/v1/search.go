package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SearchIndexReq struct {
	g.Meta `path:"/index" tags:"search" method:"get" summary:"搜索页面数据"`
}

type SearchIndexRes struct {
	g.Meta `mime:"json"`
	*model.SearchIndex
}

type SearchHelperReq struct {
	g.Meta  `path:"/helper" tags:"search" method:"get" summary:"搜索帮助"`
	Keyword string
}

type SearchHelperRes struct {
	g.Meta       `mime:"json"`
	HelpKeywords []string `json:"helpKeywords"`
}

type SearchClearHistoryReq struct {
	g.Meta `path:"/clearHistory" tags:"search" method:"post" summary:"搜索帮助"`
}

type SearchClearHistoryRes struct {
	g.Meta   `mime:"json"`
	IsDelete int `json:"is_delete"`
}
