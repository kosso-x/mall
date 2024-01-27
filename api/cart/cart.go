// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cart

import (
	"context"

	"mall/api/cart/v1"
)

type ICartV1 interface {
	CartAdd(ctx context.Context, req *v1.CartAddReq) (res *v1.CartAddRes, err error)
	CartIndex(ctx context.Context, req *v1.CartIndexReq) (res *v1.CartIndexRes, err error)
	CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (res *v1.CartUpdateRes, err error)
	CartDelete(ctx context.Context, req *v1.CartDeleteReq) (res *v1.CartDeleteRes, err error)
	CartCheck(ctx context.Context, req *v1.CartCheckReq) (res *v1.CartCheckRes, err error)
	CartGoodsCount(ctx context.Context, req *v1.CartGoodsCountReq) (res *v1.CartGoodsCountRes, err error)
	CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (res *v1.CartCheckoutRes, err error)
}
