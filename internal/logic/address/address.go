package address

import (
	"context"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/packed"
	"mall/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sAddress struct{}
)

func init() {
	service.RegisterAddress(New())
}

func New() service.IAddress {
	return &sAddress{}
}

func (s *sAddress) Detail(ctx context.Context, id int) (res *model.AddressDetail, err error) {
	dao.HiolabsAddress.Ctx(ctx).Where("user_id = ? and id = ?", cmd.CurrentUser.Id, id).Scan(&res)
	s.GetRegion(ctx, res)

	return
}

func (s *sAddress) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.
		HiolabsAddress.
		Ctx(ctx).
		Update(g.Map{
			"is_delete": 1,
		}, "user_id = ? and id = ?", cmd.CurrentUser.Id, id)

	if err != nil {
		return gerror.New("delete address error")
	}

	return
}

func (s *sAddress) Save(ctx context.Context, addr *model.AddressSave) (addr_res *model.AddressSaveRes, err error) {
	addr_res_id, err := dao.HiolabsAddress.Ctx(ctx).InsertAndGetId(g.Map{
		"name":        addr.Name,
		"mobile":      addr.Mobile,
		"province_id": addr.Province_id,
		"city_id":     addr.City_id,
		"district_id": addr.District_id,
		"address":     addr.Address,
		"is_default":  addr.Is_default,
	})

	if err != nil {
		return nil, gerror.New("insert address error")
	}

	dao.HiolabsAddress.Ctx(ctx).Where("id = ?", addr_res_id).Scan(&addr_res)

	return
}

func (s *sAddress) List(ctx context.Context) (res model.AddressList, err error) {
	err = dao.HiolabsAddress.Ctx(ctx).Where("user_id = ? and is_delete = ?", cmd.CurrentUser.Id, 0).Scan(&res)
	for i, _ := range res {
		s.GetRegion(ctx, &res[i])
	}
	if err != nil {
		return nil, gerror.New("get address list error")
	}

	return
}

func (s *sAddress) GetRegion(ctx context.Context, res *model.AddressDetail) (err error) {
	region := &model.RegionIds{
		Country_id:  res.Country_id,
		Province_id: res.Province_id,
		City_id:     res.City_id,
		District_id: res.District_id,
		Full_region: "",
	}

	values := packed.GetFullRegion(ctx, region)

	res.Province_name = values.Province_name
	res.City_name = values.City_name
	res.District_name = values.District_name
	res.Full_region = values.Full_region
	return
}
