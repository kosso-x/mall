// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"mall/api/upload/v1"
)

type IUploadV1 interface {
	UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error)
}
