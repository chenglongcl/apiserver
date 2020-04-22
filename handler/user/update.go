package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"apiserver/util"

	"github.com/gin-gonic/gin"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	menuService := userservice.User{
		ID:       r.ID,
		Password: r.Password,
		Mobile:   r.Mobile,
	}
	if errNo := menuService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
