// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"mall/api/goods/v1"
)

type IGoodsV1 interface {
	GoodsCount(ctx context.Context, req *v1.GoodsCountReq) (res *v1.GoodsCountRes, err error)
	GoodsDetail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error)
	GoodsList(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error)
	GoodsShare(ctx context.Context, req *v1.GoodsShareReq) (res *v1.GoodsShareRes, err error)
}
