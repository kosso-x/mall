// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsSearchHistory is the golang structure for table hiolabs_search_history.
type HiolabsSearchHistory struct {
	Id      uint   `json:"id"      ` //
	Keyword string `json:"keyword" ` //
	From    string `json:"from"    ` // 搜索来源，如PC、小程序、APP等
	AddTime int    `json:"addTime" ` // 搜索时间
	UserId  string `json:"userId"  ` //
}
