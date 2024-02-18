package packed

import (
	"context"
	"fmt"
	"mall/internal/dao"
	"mall/internal/model"

	"github.com/gogf/gf/util/gconv"
)

func GetFullRegion(ctx context.Context, region *model.RegionIds) (values *model.RegionValues) {
	country_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", region.Country_id).Value()
	province_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", region.Province_id).Value()
	city_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", region.City_id).Value()
	district_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", region.District_id).Value()
	full_region := fmt.Sprintf("%s%s%s%s", country_name, province_name, city_name, district_name)

	values = &model.RegionValues{
		Country_name:  gconv.String(country_name),
		Province_name: gconv.String(province_name),
		City_name:     gconv.String(city_name),
		District_name: gconv.String(district_name),
		Full_region:   full_region,
	}
	return
}
