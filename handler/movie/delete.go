package movie

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/movieservice"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	movieService := &movieservice.Movie{
		ID: bson.ObjectIdHex(r.ID),
	}
	if errNo := movieService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
