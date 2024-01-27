// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsAddress is the golang structure for table hiolabs_address.
type HiolabsAddress struct {
	Id         uint   `json:"id"         ` //
	Name       string `json:"name"       ` //
	UserId     uint   `json:"userId"     ` //
	CountryId  int    `json:"countryId"  ` //
	ProvinceId int    `json:"provinceId" ` //
	CityId     int    `json:"cityId"     ` //
	DistrictId int    `json:"districtId" ` //
	Address    string `json:"address"    ` //
	Mobile     string `json:"mobile"     ` //
	IsDefault  uint   `json:"isDefault"  ` //
	IsDelete   int    `json:"isDelete"   ` //
}
