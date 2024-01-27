// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsShipper is the golang structure for table hiolabs_shipper.
type HiolabsShipper struct {
	Id           int    `json:"id"           ` //
	Name         string `json:"name"         ` // 快递公司名称
	Code         string `json:"code"         ` // 快递公司代码
	SortOrder    int    `json:"sortOrder"    ` // 排序
	MonthCode    string `json:"monthCode"    ` //
	CustomerName string `json:"customerName" ` //
	Enabled      int    `json:"enabled"      ` //
}
