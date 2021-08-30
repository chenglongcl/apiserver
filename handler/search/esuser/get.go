package esuser

import (
	"apiserver/elastic/eswire"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func MGet(c *gin.Context) {
	var (
		r MGetRequest
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ids := make([]uint64, 0)
	for _, id := range strings.Split(r.IDS, ",") {
		d, _ := strconv.Atoi(id)
		ids = append(ids, uint64(d))
	}

	userService := eswire.InitUserService()
	users, err := userService.MGet(c, ids)
	if err != nil {
		handler.SendResponse(c, errno.ErrGetDocument, nil)
		return
	}
	handler.SendResponse(c, nil, users)
}
