package upload

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/uploadservice"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	uploadService := uploadservice.Upload{}
	path, fileName, errNo := uploadService.UploadFile(file)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, UploadResponse{
		Url:      viper.GetString("file_domain") + "/" + path,
		Path:     path,
		FileName: fileName,
	})
}
