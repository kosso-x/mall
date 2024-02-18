// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mime/multipart"
	"mall/internal/model"
)

type (
	IUpload interface{
		UploadAvatar(ctx context.Context, files map[string][]*multipart.FileHeader) (file *model.FileInfo, err error)
	}
)

var (
	localUpload IUpload
)

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
