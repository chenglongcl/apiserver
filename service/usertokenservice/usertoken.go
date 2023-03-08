package usertokenservice

import (
	"apiserver/dal/apiserverdb/apiservermodel"
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/pkg/errno"
	"apiserver/pkg/gormx"
	"apiserver/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"time"
)

type userToken struct {
	UserID         uint64
	Token          string
	ExpireTime     time.Time
	RefreshTime    time.Time
	serviceOptions *service.Options
	ctx            *gin.Context
}

type UserToken = *userToken

// NewUserTokenService
// @Description:
// @param ctx
// @param opts
// @return UserToken
func NewUserTokenService(ctx *gin.Context, opts ...service.Option) UserToken {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &userToken{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

// GetByUserID
// @Description:
// @receiver a
// @return *apiservermodel.TbUserToken
// @return *errno.Errno
func (a UserToken) GetByUserID() (*apiservermodel.TbUserToken, *errno.Errno) {
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	userToken, err := qc.TbUserToken.Where(q.TbUserToken.UserID.Eq(a.UserID)).Take()
	return userToken, gormx.HandleError(err)
}

// RecordToken
// @Description:
// @receiver a
// @return *errno.Errno
func (a UserToken) RecordToken() *errno.Errno {
	err := apiserverquery.Q.WithContext(a.ctx).TbUserToken.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"user_id":      a.UserID,
			"token":        a.Token,
			"expire_time":  a.ExpireTime,
			"refresh_time": a.RefreshTime,
			"updated_at":   time.Now(),
		}),
	}).Create(&apiservermodel.TbUserToken{
		UserID:      a.UserID,
		Token:       a.Token,
		ExpireTime:  &a.ExpireTime,
		RefreshTime: &a.RefreshTime,
	})
	return gormx.HandleError(err)
}
