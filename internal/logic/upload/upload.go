package upload

import (
	"context"
	"fmt"
	"io"
	"mall/internal/model"
	"mall/internal/service"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type (
	sUpload struct{}
)

func init() {
	service.RegisterUpload(New())
}

func New() service.IUpload {
	return &sUpload{}
}

func (s *sUpload) UploadAvatar(ctx context.Context, files map[string][]*multipart.FileHeader) (fileRes *model.FileInfo, err error) {
	var upload_files *model.UploadFiles
	if err = gconv.Struct(files, &upload_files); err != nil {
		return
	}

	fileRes, err = s.SaveAvatar(ctx, &upload_files.UploadFiles[0], files["upload_file"][0])
	if err != nil {
		return nil, err
	}

	return
}

func (s *sUpload) SaveAvatar(ctx context.Context, u_file *model.UFile, u_file_header *multipart.FileHeader) (fileRes *model.FileInfo, err error) {
	var (
		fileInfo model.FileInfo
		content  []byte
		tmp      = make([]byte, 128)
	)

	file_types := u_file.Header.ContentType[0]
	file_type_text := strings.Split(gconv.String(file_types), "/")[1]
	fileInfo.Name = uuid.New().String() + "." + file_type_text
	file, err := u_file_header.Open()

	for {
		n, err := file.Read(tmp) //每次读取128个字节
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		//每次读取的内容都追加到已知的byte切片中
		content = append(content, tmp[:n]...)
	}

	config := g.Cfg("mall")
	static_path, err := config.Get(ctx, "default.static_path")
	if err != nil {
		return
	}
	avatar_path, err := config.Get(ctx, "default.avatar_path")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s%s/%s", static_path, avatar_path, fileInfo.Name)
	new_file, err := os.Create(path)
	if err != nil {
		return
	}
	defer new_file.Close()
	defer file.Close()
	new_file.Write(content)
	fileInfo.FileUrl = fmt.Sprintf("%s/%s", avatar_path, fileInfo.Name)

	fileRes = &fileInfo
	return
}
