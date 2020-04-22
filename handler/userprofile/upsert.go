package userprofile

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/userprofileservice"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"time"
)

func Upsert(c *gin.Context) {
	var r UpsertRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	userID, _ := c.Get("userID")
	dateOfBirth, _ := time.ParseInLocation("2006-01-02", r.DateOfBirth, time.Local)
	userProfileService := userprofileservice.UserProfile{
		UserID:      userID.(uint64),
		RealName:    r.RealName,
		Sex:         r.Sex,
		DateOfBirth: dateOfBirth,
	}
	if errNo := userProfileService.Upsert(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
