// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package category

import (
	"context"

	"mall/api/category/v1"
)

type ICategoryV1 interface {
	Categories(ctx context.Context, req *v1.CategoriesReq) (res *v1.CategoriesRes, err error)
	CategoryCurrent(ctx context.Context, req *v1.CategoryCurrentReq) (res *v1.CategoryCurrentRes, err error)
	CategoryCurrentList(ctx context.Context, req *v1.CategoryCurrentListReq) (res *v1.CategoryCurrentListRes, err error)
}
