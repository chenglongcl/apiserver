package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

// Create
// @Description:
// @param c
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.Username = r.Username
	userService.Password = r.Password
	userService.Mobile = r.Mobile
	user, errNo := userService.Add()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	t, e, re, _ := token.Sign(c, token.Context{ID: user.ID, Username: user.Username}, "")
	rep := CreateResponse{
		Username:         r.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	}
	handler.SendResponse(c, nil, rep)
}
