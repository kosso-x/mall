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
	ISettings interface{
		Show(ctx context.Context) (res *model.ShowSettings, err error)
		Save(ctx context.Context, settings *model.SaveSettings) (err error)
		UserDetail(ctx context.Context) (res *model.UserDetail, err error)
	}
)

var (
	localSettings ISettings
)

func Settings() ISettings {
	if localSettings == nil {
		panic("implement not found for interface ISettings, forgot register?")
	}
	return localSettings
}

func RegisterSettings(i ISettings) {
	localSettings = i
}
