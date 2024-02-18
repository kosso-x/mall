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
	IFootprint interface{
		List(ctx context.Context, page, size int) (res *model.FootprintList, err error)
		Delete(ctx context.Context, fpId int) (err error)
	}
)

var (
	localFootprint IFootprint
)

func Footprint() IFootprint {
	if localFootprint == nil {
		panic("implement not found for interface IFootprint, forgot register?")
	}
	return localFootprint
}

func RegisterFootprint(i IFootprint) {
	localFootprint = i
}
