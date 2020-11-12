package oss

import (
	"apiserver/pkg/oss/client"
	"github.com/gin-gonic/gin"
)

import (
	. "apiserver/handler"
	"net/http"
)

func WebUploadSign(c *gin.Context) {
	sign, errNo := client.DefaultAliClient().WebUploadSign()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	c.JSON(http.StatusOK, sign)
}
