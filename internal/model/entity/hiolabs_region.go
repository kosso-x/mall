// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsRegion is the golang structure for table hiolabs_region.
type HiolabsRegion struct {
	Id       uint   `json:"id"       ` //
	ParentId uint   `json:"parentId" ` //
	Name     string `json:"name"     ` //
	Type     int    `json:"type"     ` //
	AgencyId uint   `json:"agencyId" ` //
	Area     uint   `json:"area"     ` // 方位，根据这个定运费
	AreaCode string `json:"areaCode" ` // 方位代码
	FarArea  uint   `json:"farArea"  ` // 偏远地区
}
