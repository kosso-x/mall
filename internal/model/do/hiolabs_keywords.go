// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsKeywords is the golang structure of table hiolabs_keywords for DAO operations like Where/Data.
type HiolabsKeywords struct {
	g.Meta    `orm:"table:hiolabs_keywords, do:true"`
	Keyword   interface{} //
	IsHot     interface{} //
	IsDefault interface{} //
	IsShow    interface{} //
	SortOrder interface{} //
	SchemeUrl interface{} // 关键词的跳转链接
	Id        interface{} //
	Type      interface{} //
}
