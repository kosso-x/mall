// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package address

import (
	"context"

	"mall/api/address/v1"
)

type IAddressV1 interface {
	AddressDetail(ctx context.Context, req *v1.AddressDetailReq) (res *v1.AddressDetailRes, err error)
	AddressDelete(ctx context.Context, req *v1.AddressDeleteReq) (res *v1.AddressDeleteRes, err error)
	AddressSave(ctx context.Context, req *v1.AddressSaveReq) (res *v1.AddressSaveRes, err error)
	AddressGet(ctx context.Context, req *v1.AddressGetReq) (res *v1.AddressGetRes, err error)
}
