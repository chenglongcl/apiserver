package esuser

import (
	"apiserver/elastic/esmodel"
	"apiserver/elastic/eswire"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var err error
	user := &esmodel.UserEs{}
	if err := c.Bind(&user); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	users := make([]*esmodel.UserEs, 0)
	users = append(users, user)
	userService := eswire.InitUserService()
	if user, err = userService.Get(c, user.ID); err != nil {
		handler.SendResponse(c, errno.ErrUpdateDocument, nil)
		return
	}
	if user == nil {
		handler.SendResponse(c, errno.ErrDocumentNotFound, nil)
		return
	}
	if err := userService.BatchUpdate(c, users); err != nil {
		handler.SendResponse(c, errno.ErrUpdateDocument, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
