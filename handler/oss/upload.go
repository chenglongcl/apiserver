package oss

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/oss"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	var r UploadOssRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	switch r.OssName {
	case "aliYunOss":
		fileUrl, errNo := oss.SelectClient("ali").Upload(file, header)
		if errNo != nil {
			SendResponse(c, errNo, nil)
			return
		}
		SendResponse(c, nil, fileUrl)
	}
}
