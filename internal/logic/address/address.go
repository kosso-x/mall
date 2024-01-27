package address

import (
	"context"
	"fmt"
	"mall/internal/cmd"
	"mall/internal/dao"
	"mall/internal/model"
	"mall/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
	province_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", res.Province_id).Value()
	city_name, _ := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", res.City_id).Value()
	district_name, err := dao.HiolabsRegion.Ctx(ctx).Fields("name").Where("id = ?", res.District_id).Value()

	if err != nil {
		return gerror.New("get address region error")
	}

	res.Province_name = gconv.String(province_name)
	res.City_name = gconv.String(city_name)
	res.District_name = gconv.String(district_name)
	res.Full_region = fmt.Sprintf("%s%s%s", province_name, city_name, district_name)
	return
}
