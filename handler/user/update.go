package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userID, _ := c.Get("userID")
	menuService := userservice.User{
		ID:       userID.(uint64),
		Password: r.Password,
		Mobile:   r.Mobile,
	}
	if errNo := menuService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
