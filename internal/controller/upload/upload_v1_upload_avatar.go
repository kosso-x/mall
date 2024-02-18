package upload

import (
	"context"

	v1 "mall/api/upload/v1"
	"mall/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error) {
	files := g.RequestFromCtx(ctx).GetMultipartForm().File
	r, err := service.Upload().UploadAvatar(ctx, files)
	if err != nil {
		return
	}

	res = &v1.UploadAvatarRes{
		FileInfo: r,
	}

	return
}
