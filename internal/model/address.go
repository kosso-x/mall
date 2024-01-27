package model

import "mall/internal/model/entity"

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
	Name        string
	Mobile      string
	Province_id int
	City_id     int
	District_id int
	Address     string
	Is_default  int
}

type (
	AddressSaveRes entity.HiolabsAddress
)

// get
type (
	AddressList []AddressDetail
)
