package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
)

// Get
// @Description:
// @param c
func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.ID = r.ID
	user, errNo := userService.GetByID()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	if user.ID == 0 {
		SendResponse(c, errno.ErrNotUserExist, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		ID:         user.ID,
		Username:   user.Username,
		Mobile:     user.Mobile,
		CreateTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
