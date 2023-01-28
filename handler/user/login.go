package user

import (
	"apiserver/handler"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"apiserver/service/userservice"
	"apiserver/service/usertokenservice"
	"github.com/gin-gonic/gin"
	"time"
)

// Login
// @Description:
// @param c

func Login(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.Username = r.Username
	user, errNo := userService.GetByUsername()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if user.ID == 0 {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	//Compare the login password with user password
	if err := auth.Compare(user.Password, r.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	// Sign the json web token.
	t, e, re, err := token.Sign(c, token.Context{ID: user.ID, Username: user.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}
	go func() {
		expireTime, _ := time.ParseInLocation("2006-01-02 15:04:05", e, time.Local)
		RefreshTime, _ := time.ParseInLocation("2006-01-02 15:04:05", re, time.Local)
		userTokenService := usertokenservice.NewUserTokenService(c)
		userTokenService.UserID = user.ID
		userTokenService.Token = t
		userTokenService.ExpireTime = expireTime
		userTokenService.RefreshTime = RefreshTime
		_ = userTokenService.RecordToken()
	}()
	handler.SendResponse(c, nil, CreateResponse{
		Username:         user.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	})
}

// Logout
// @Description:
// @param c

func Logout(c *gin.Context) {
	userID := c.GetUint64("userID")
	userTokenService := usertokenservice.NewUserTokenService(c)
	userTokenService.UserID = userID
	userToken, errNo := userTokenService.GetByUserID()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if userToken.UserID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	go func() {
		tokenCtx := &token.Context{
			ID:         userID,
			ExpiredAt:  userToken.ExpireTime.Unix(),
			RefreshExp: userToken.RefreshTime.Unix(),
		}
		token.BLackListToken(userToken.Token, tokenCtx)
	}()
	handler.SendResponse(c, nil, nil)
}
