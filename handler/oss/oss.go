package oss

type UploadOssRequest struct {
	OssName string `form:"ossName" binding:"required"`
}
