// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsFreightTemplateGroup is the golang structure for table hiolabs_freight_template_group.
type HiolabsFreightTemplateGroup struct {
	Id           uint    `json:"id"           ` //
	TemplateId   int     `json:"templateId"   ` //
	IsDefault    int     `json:"isDefault"    ` // 默认，area:0
	Area         string  `json:"area"         ` // 0位默认，
	Start        int     `json:"start"        ` //
	StartFee     float64 `json:"startFee"     ` //
	Add          int     `json:"add"          ` //
	AddFee       float64 `json:"addFee"       ` //
	FreeByNumber int     `json:"freeByNumber" ` // 0没有设置
	FreeByMoney  float64 `json:"freeByMoney"  ` // 0没设置
	IsDelete     int     `json:"isDelete"     ` //
}
