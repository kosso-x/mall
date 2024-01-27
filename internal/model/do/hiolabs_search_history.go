// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// HiolabsSearchHistory is the golang structure of table hiolabs_search_history for DAO operations like Where/Data.
type HiolabsSearchHistory struct {
	g.Meta  `orm:"table:hiolabs_search_history, do:true"`
	Id      interface{} //
	Keyword interface{} //
	From    interface{} // 搜索来源，如PC、小程序、APP等
	AddTime interface{} // 搜索时间
	UserId  interface{} //
}
