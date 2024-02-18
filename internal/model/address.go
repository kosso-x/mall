package model

// detail
type AddressDetail struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	User_id       int    `json:"user_id"`
	Country_id    int    `json:"country_id"`
	Province_id   int    `json:"province_id"`
	City_id       int    `json:"city_id"`
	District_id   int    `json:"district_id"`
	Address       string `json:"address"`
	Mobile        string `json:"mobile"`
	Is_default    int    `json:"is_default"`
	Is_delete     int    `json:"is_delete"`
	Province_name string `json:"province_name"`
	City_name     string `json:"city_name"`
	District_name string `json:"district_name"`
	Full_region   string `json:"full_region"`
}

// save
type AddressSave struct {
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	Province_id int    `json:"province_id"`
	City_id     int    `json:"city_id"`
	District_id int    `json:"district_id"`
	Address     string `json:"address"`
	Is_default  int    `json:"is_default"`
}

type (
	AddressSaveRes HAddress
)

type HAddress struct {
	Id         uint   `json:"id"         `  //
	Name       string `json:"name"       `  //
	UserId     uint   `json:"user_id"     ` //
	CountryId  int    `json:"country_id"  ` //
	ProvinceId int    `json:"province_id" ` //
	CityId     int    `json:"city_id"     ` //
	DistrictId int    `json:"district_id" ` //
	Address    string `json:"address"    `  //
	Mobile     string `json:"mobile"     `  //
	IsDefault  uint   `json:"is_default"  ` //
	IsDelete   int    `json:"is_delete"   ` //
}

// get
type (
	AddressList []AddressDetail
)
