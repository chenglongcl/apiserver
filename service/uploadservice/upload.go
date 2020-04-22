package uploadservice

import (
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Upload struct {
}

func (a *Upload) UploadFile(file *multipart.FileHeader) (string, string, *errno.Errno) {
	// Get the file suffix name
	fileSuffix := strings.ToLower(path.Ext(file.Filename))
	// Rename filename and set savePath
	savePath := viper.GetString("upload_path")
	// Set Folder by date
	date := time.Now().Format("20060102")
	//Folder isNotExist
	if _, err := os.Stat(savePath + date); err != nil && os.IsNotExist(err) {
		os.MkdirAll(savePath+date, os.ModePerm)
	}
	//Set saveFileName
	saveFileName, _ := util.GenStr(16)
	dst := savePath + date + "/" + saveFileName + fileSuffix
	c := &gin.Context{}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", "", errno.ErrUploadFail
	}
	return dst, saveFileName + fileSuffix, nil
}
