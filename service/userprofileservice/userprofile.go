package userprofileservice

import (
	"apiserver/mgomodel"
	"apiserver/pkg/errno"
	"time"
)

type UserProfile struct {
	UserID      uint64
	RealName    string
	Sex         int64
	DateOfBirth time.Time
}

func (a *UserProfile) Upsert() *errno.Errno {
	if err := mgomodel.UpsertUserProfile(map[string]interface{}{
		"user_id":       a.UserID,
		"real_name":     a.RealName,
		"sex":           a.Sex,
		"date_of_birth": a.DateOfBirth,
	}); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
