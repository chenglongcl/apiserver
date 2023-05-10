package user

import (
	"apiserver/handler"
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
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.ID = r.ID
	user, errNo := userService.GetByID()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if user == nil || user.ID == 0 {
		handler.SendResponse(c, errno.ErrNotUserExist, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:         user.ID,
		Username:   user.Username,
		Mobile:     user.Mobile,
		CreateTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
