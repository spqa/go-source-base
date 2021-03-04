package media

import (
	"io"
	"mcm-api/pkg/common"
)

type UploadType string

const (
	Document UploadType = "document"
	Image    UploadType = "image"
)

type UploadResult struct {
	Key string
}

type FileUploadOriginalReq struct {
	File io.ReadSeeker
	Size int64
	Name string
	User *common.LoggedInUser
}

type FileUploadPreviewReq struct {
	File io.Reader
	Name string
	User common.LoggedInUser
}

type UploadQuery struct {
	Type UploadType `query:"type" enums:"document,image"`
}
