package handler

import (
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RegularListResponse struct {
	util.Page
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendResponseUnauthorized(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusUnauthorized, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
