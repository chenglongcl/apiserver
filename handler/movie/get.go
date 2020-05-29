package movie

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/movieservice"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func Get(c *gin.Context) {
	var r GetRequest
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
	movie, errNo := movieService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	if movie.ID == "" {
		SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		ID:          movie.ID.Hex(),
		MovieName:   movie.MovieName,
		Description: movie.Description,
		Thumb:       movie.Thumb,
		ReleaseTime: movie.ReleaseTime.Local().Format("2006-01-02 15:04:05"),
		BoxOffice:   movie.BoxOffice,
	})
}
