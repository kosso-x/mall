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
	ISearch interface{
		Index(ctx context.Context) (res *model.SearchIndex, err error)
		Helper(ctx context.Context, keyword string) (texts []string, err error)
		ClearHistory(ctx context.Context) (err error)
	}
)

var (
	localSearch ISearch
)

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for interface ISearch, forgot register?")
	}
	return localSearch
}

func RegisterSearch(i ISearch) {
	localSearch = i
}
