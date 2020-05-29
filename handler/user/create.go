package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.User{
		Username: r.Username,
		Password: r.Password,
		Mobile:   r.Mobile,
	}
	id, errNo := userService.Add()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	t, e, re, _ := token.Sign(c, token.Context{ID: id, Username: userService.Username}, "")
	rep := CreateResponse{
		Username:         r.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	}
	SendResponse(c, nil, rep)
}
