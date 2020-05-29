package movie

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/movieservice"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			SendResponse(c, errno.ErrObjectIdHex, nil)
			return
		}
	}()
	movieService := &movieservice.Movie{
		ID: bson.ObjectIdHex(r.ID),
	}
	if errNo := movieService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
