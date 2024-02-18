package v1

import (
	"mall/internal/model"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
)

type UploadAvatarReq struct {
	g.Meta     `path:"/uploadAvatar" tags:"upload" method:"post" summary:"上传头像"`
	UploadFile *ghttp.UploadFile
}

type UploadAvatarRes struct {
	g.Meta `mime:"json"`
	*model.FileInfo
}
