package common

import (
	"apiserver/pkg/errno"
	"mime/multipart"
)

type OSSClient interface {
	Upload(file multipart.File, header *multipart.FileHeader) (string, *errno.Errno)
}
