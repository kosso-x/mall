package model

import "github.com/gogf/gf/frame/g"

type UploadFiles struct {
	UploadFiles []UFile `json:"upload_file"`
}

type UFile struct {
	Filename string             `json:"Filename"`
	Header   ContentDisposition `json:"Header"`
	Size     int                `json:"Size"`
}

type ContentDisposition struct {
	ContentDisposition []*g.Var `json:"Content-Disposition"`
	ContentType        []*g.Var `json:"Content-Type"`
}

type FileInfo struct {
	Name    string `json:"name"`
	FileUrl string `json:"fileUrl"`
}
