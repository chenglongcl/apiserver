package movie

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/movieservice"
	"github.com/gin-gonic/gin"
	"time"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	releaseTime, _ := time.ParseInLocation("2006-01-02 15:04:05", r.ReleaseTime, time.Local)
	movieService := &movieservice.Movie{
		MovieName:   r.MovieName,
		Description: r.Description,
		Thumb:       r.Thumb,
		ReleaseTime: releaseTime,
		BoxOffice:   r.BoxOffice,
	}
	if errNo := movieService.Add(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
