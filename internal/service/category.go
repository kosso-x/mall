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
	ICategory interface {
		List(ctx context.Context) (list model.CategoryList, err error)
		One(ctx context.Context, id int) (current model.CurrentCategory, err error)
		CurrentList(ctx context.Context, Id, Page, Size int) (res model.CategoryCurrentList, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
