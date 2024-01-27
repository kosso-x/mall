// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsFreightTemplate is the golang structure for table hiolabs_freight_template.
type HiolabsFreightTemplate struct {
	Id           uint    `json:"id"           ` //
	Name         string  `json:"name"         ` // 运费模板名称
	PackagePrice float64 `json:"packagePrice" ` // 包装费用
	FreightType  int     `json:"freightType"  ` // 0按件，1按重量
	IsDelete     int     `json:"isDelete"     ` //
}
