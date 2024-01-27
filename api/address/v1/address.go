package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddressDetailReq struct {
	g.Meta `path:"/addressDetail" tags:"address" method:"get" summary:"收货地址详情"`
	Id     int
}

type AddressDetailRes struct {
	g.Meta `mime:"json"`
	*model.AddressDetail
}

type AddressDeleteReq struct {
	g.Meta `path:"/deleteAddress" tags:"address" method:"post" summary:"删除收货地址"`
	Id     int
}

type AddressDeleteRes struct {
	g.Meta    `mime:"json"`
	Is_delete int `json:"is_delete"`
}

type AddressSaveReq struct {
	g.Meta `path:"/saveAddress" tags:"address" method:"post" summary:"保存收货地址"`
	*model.AddressSave
}

type AddressSaveRes struct {
	g.Meta `mime:"json"`
	*model.AddressSaveRes
}

type AddressGetReq struct {
	g.Meta `path:"/getAddresses" tags:"address" method:"get" summary:"收货地址"`
}

type AddressGetRes struct {
	g.Meta             `mime:"json"`
	*model.AddressList `json:"addressList"`
}
