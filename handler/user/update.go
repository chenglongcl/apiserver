package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

// Update
// @Description:
// @param c

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userID := c.GetUint64("userID")
	userService := userservice.NewUserService(c)
	userService.ID = userID
	userService.Password = r.Password
	userService.Mobile = r.Mobile
	if errNo := userService.Edit(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
