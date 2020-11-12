package usertokenservice

import (
	"apiserver/model"
	"apiserver/pkg/errno"
	"time"
)

type UserToken struct {
	UserID      uint64
	Token       string
	ExpireTime  time.Time
	RefreshTime time.Time
}

func (a *UserToken) RecordToken() *errno.Errno {
	userToken, err := model.GetUserToken(a.UserID)
	if err != nil {
		return errno.ErrDatabase
	}
	data := map[string]interface{}{
		"user_id":      a.UserID,
		"token":        a.Token,
		"expire_time":  a.ExpireTime,
		"refresh_time": a.RefreshTime,
	}
	if userToken.UserID > 0 {
		_ = model.EditUserToken(data)
	} else {
		_ = model.AddUserToken(data)
	}
	return nil
}

func (a *UserToken) Get() (*model.UserToken, *errno.Errno) {
	userToken, err := model.GetUserToken(a.UserID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return userToken, nil
}
