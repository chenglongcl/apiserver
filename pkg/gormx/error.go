package gormx

import (
	"apiserver/pkg/errno"
	"github.com/chenglongcl/log"
	"gorm.io/gorm"
)

func HandleError(err error) *errno.Errno {
	if err == gorm.ErrRecordNotFound {
		return errno.ErrRecordNotFound
	} else if err != nil {
		log.Error("database error", err)
		return errno.ErrDatabase
	}
	return nil
}
