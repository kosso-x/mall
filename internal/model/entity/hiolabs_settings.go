// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// HiolabsSettings is the golang structure for table hiolabs_settings.
type HiolabsSettings struct {
	Id                 int    `json:"id"                 ` //
	AutoDelivery       int    `json:"autoDelivery"       ` //
	Name               string `json:"name"               ` //
	Tel                string `json:"tel"                ` //
	ProvinceName       string `json:"provinceName"       ` //
	CityName           string `json:"cityName"           ` //
	ExpAreaName        string `json:"expAreaName"        ` //
	Address            string `json:"address"            ` //
	DiscoveryImgHeight int    `json:"discoveryImgHeight" ` //
	DiscoveryImg       string `json:"discoveryImg"       ` //
	GoodsId            int    `json:"goodsId"            ` //
	CityId             int    `json:"cityId"             ` //
	ProvinceId         int    `json:"provinceId"         ` //
	DistrictId         int    `json:"districtId"         ` //
	Countdown          int    `json:"countdown"          ` // 10分钟倒计时
	Reset              int    `json:"reset"              ` //
}
