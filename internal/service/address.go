// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mall/internal/model"
)

type (
	IAddress interface{
		Detail(ctx context.Context, id int) (res *model.AddressDetail, err error)
		Delete(ctx context.Context, id int) (err error)
		Save(ctx context.Context, address *model.AddressSave) (addr_res *model.AddressSaveRes, err error)
		List(ctx context.Context) (res model.AddressList, err error)
		GetRegion(ctx context.Context, res *model.AddressDetail) (err error)
	}
)

var (
	localAddress IAddress
)

func Address() IAddress {
	if localAddress == nil {
		panic("implement not found for interface IAddress, forgot register?")
	}
	return localAddress
}

func RegisterAddress(i IAddress) {
	localAddress = i
}
