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
	IRegion interface{
		List(ctx context.Context, parentId int) (res model.RegionList, err error)
	}
)

var (
	localRegion IRegion
)

func Region() IRegion {
	if localRegion == nil {
		panic("implement not found for interface IRegion, forgot register?")
	}
	return localRegion
}

func RegisterRegion(i IRegion) {
	localRegion = i
}
