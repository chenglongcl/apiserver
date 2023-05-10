package gormx

import (
	"apiserver/pkg/errno"
	"github.com/chenglongcl/log"
	"gorm.io/gorm"
)

func HandleError(err error) *errno.Errno {
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("database error", err)
		return errno.ErrDatabase
	}
	return nil
}
