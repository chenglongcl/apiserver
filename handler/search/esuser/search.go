package esuser

import (
	"apiserver/elastic/esmodel"
	"apiserver/elastic/eswire"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	req := esmodel.SearchUserRequest{}
	if err := c.Bind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := eswire.InitUserService()
	users, count, err := userService.Search(c, &req)
	if err != nil {
		handler.SendResponse(c, errno.ErrSearchDocument, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, uint64(req.Page), uint64(req.Limit), users))
}
