package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

// Delete
// @Description:
// @param c

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.ID = r.ID
	if errNo := userService.Delete(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
