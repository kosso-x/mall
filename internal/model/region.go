package model

import "mall/internal/model/entity"

// list
type (
	RegionList []entity.HiolabsRegion
)

type RegionIds struct {
	Country_id  int    `json:"country_id"`
	Province_id int    `json:"province_id"`
	City_id     int    `json:"city_id"`
	District_id int    `json:"district_id"`
	Full_region string `json:"full_region"`
}

type RegionValues struct {
	Country_name  string `json:"country_name"`
	Province_name string `json:"province_name"`
	City_name     string `json:"city_name"`
	District_name string `json:"district_name"`
	Full_region   string `json:"full_region"`
}
