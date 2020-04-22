package handler

import (
	"github.com/gin-gonic/gin"
	"apiserver/pkg/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	TotalCount uint64        `json:"total_count"`
	PageSize   uint64        `json:"page_size"`
	List       []interface{} `json:"list"`
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
