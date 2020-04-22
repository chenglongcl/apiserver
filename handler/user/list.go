package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"apiserver/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var (
		r  ListRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	userService := userservice.User{
		Username: r.UserName,
	}
	info, count, errNo := userService.GetUserList(ps)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
