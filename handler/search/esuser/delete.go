package esuser

import (
	"apiserver/elastic/esmodel"
	"apiserver/elastic/eswire"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	user := &esmodel.UserEs{}
	if err := c.Bind(&user); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	users := make([]*esmodel.UserEs, 0)
	users = append(users, user)
	userService := eswire.InitUserService()
	if err := userService.BatchDel(c, users); err != nil {
		handler.SendResponse(c, errno.ErrDeleteDocument, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
