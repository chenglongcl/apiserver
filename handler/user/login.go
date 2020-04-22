package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"apiserver/service/usertokenservice"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	//Compare the login password with user password
	if err := auth.Compare(user.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	// Sign the json web token.
	t, e, re, err := token.Sign(c, token.Context{ID: user.ID, Username: user.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	go func() {
		expireTime, _ := time.ParseInLocation("2006-01-02 15:04:05", e, time.Local)
		RefreshTime, _ := time.ParseInLocation("2006-01-02 15:04:05", re, time.Local)
		userTokenService := &usertokenservice.UserToken{
			UserID:      user.ID,
			Token:       t,
			ExpireTime:  expireTime,
			RefreshTime: RefreshTime,
		}
		_ = userTokenService.RecordToken()
	}()
	SendResponse(c, nil, CreateResponse{
		Username:         user.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	})
}

//用户登出
func Logout(c *gin.Context) {
	userID, _ := c.Get("userID")
	userTokenService := &usertokenservice.UserToken{
		UserID: userID.(uint64),
	}
	userToken, errNo := userTokenService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	if userToken.UserID == 0 {
		SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	go func() {
		tokenCtx := &token.Context{
			ID:         userID.(uint64),
			ExpiredAt:  userToken.ExpireTime.Unix(),
			RefreshExp: userToken.RefreshTime.Unix(),
		}
		token.BLackListToken(userToken.Token, tokenCtx)
	}()
	SendResponse(c, nil, nil)
}
