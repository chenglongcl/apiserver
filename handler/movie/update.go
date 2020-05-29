package movie

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/movieservice"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"time"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	releaseTime, _ := time.ParseInLocation("2006-01-02 15:04:05", r.ReleaseTime, time.Local)
	defer func() {
		if err := recover(); err != nil {
			SendResponse(c, errno.ErrObjectIdHex, nil)
			return
		}
	}()
	movieService := &movieservice.Movie{
		ID:          bson.ObjectIdHex(r.ID),
		MovieName:   r.MovieName,
		Description: r.Description,
		Thumb:       r.Thumb,
		ReleaseTime: releaseTime,
		BoxOffice:   r.BoxOffice,
	}
	if errNo := movieService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
