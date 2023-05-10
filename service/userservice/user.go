package userservice

import (
	"apiserver/dal/apiserverdb/apiserverentity"
	"apiserver/dal/apiserverdb/apiservermodel"
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/gormx"
	"apiserver/service"
	"apiserver/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type user struct {
	ID             uint64
	Username       string
	Password       string
	Mobile         string
	serviceOptions *service.Options
	ctx            *gin.Context
}

type User = *user

// NewUserService
// @Description:
// @param ctx
// @param opts
// @return *User
func NewUserService(ctx *gin.Context, opts ...service.Option) User {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &user{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

// Add
// @Description:
// @receiver a
// @return *apiservermodel.TbUser
// @return *errno.Errno
func (a User) Add() (*apiservermodel.TbUser, *errno.Errno) {
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	var (
		userModel *apiservermodel.TbUser
		err       error
	)
	if count, _ := qc.TbUser.Where(q.TbUser.Username.Eq(a.Username)).Unscoped().Count(); count > 0 {
		return userModel, errno.ErrUserExist
	}
	password, _ := auth.Encrypt(a.Password)
	userModel = &apiservermodel.TbUser{
		Username: a.Username,
		Password: password,
		Mobile:   a.Mobile,
		Status:   true,
	}
	err = qc.TbUser.Create(userModel)
	return userModel, gormx.HandleError(err)
}

// Get
// @Description:
// @receiver a
// @param fields
// @param conditions
// @return *apiservermodel.TbUser
// @return *errno.Errno
func (a User) Get(fields []field.Expr, conditions []gen.Condition) (*apiservermodel.TbUser, *errno.Errno) {
	userModel, err := apiserverquery.Q.WithContext(a.ctx).TbUser.Select(fields...).Where(conditions...).Take()
	return userModel, gormx.HandleError(err)
}

// GetByID
// @Description:
// @receiver a
// @return *apiservermodel.TbUser
// @return *errno.Errno
func (a User) GetByID() (*apiservermodel.TbUser, *errno.Errno) {
	userModel, err := apiserverquery.Q.WithContext(a.ctx).TbUser.Where(
		apiserverquery.Q.TbUser.ID.Eq(a.ID),
	).Take()
	return userModel, gormx.HandleError(err)
}

// GetByUsername
// @Description:
// @receiver a
// @return *apiservermodel.TbUser
// @return *errno.Errno
func (a User) GetByUsername() (*apiservermodel.TbUser, *errno.Errno) {
	userModel, err := apiserverquery.Q.WithContext(a.ctx).TbUser.Where(
		apiserverquery.Q.TbUser.Username.Eq(a.Username),
	).Take()
	return userModel, gormx.HandleError(err)
}

// Edit
// @Description:
// @receiver a
// @return *errno.Errno
func (a User) Edit() *errno.Errno {
	var password string
	if a.Password != "" {
		password, _ = auth.Encrypt(a.Password)
	}
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	if count, _ := qc.TbUser.Where(q.TbUser.ID.Eq(a.ID)).Count(); count == 0 {
		return errno.ErrUserNotFound
	}
	baseQuery := qc.TbUser.Where(q.TbUser.ID.Eq(a.ID))
	if password == "" {
		baseQuery = baseQuery.Omit(q.TbUser.Password)
	}
	_, err := baseQuery.Updates(map[string]interface{}{
		"password": password,
		"mobile":   a.Mobile,
	})
	return gormx.HandleError(err)
}

// Delete
// @Description:
// @receiver a
// @return *errno.Errno
func (a User) Delete() *errno.Errno {
	_, err := apiserverquery.Q.WithContext(a.ctx).TbUser.Where(
		apiserverquery.Q.TbUser.ID.Eq(a.ID),
	).Delete()
	return gormx.HandleError(err)
}

// InfoList
// @Description:
// @receiver a
// @param listParams
// @return []*apiserverentity.UserInfo
// @return uint64
// @return *errno.Errno
func (a User) InfoList(listParams *service.ListParams) ([]*apiserverentity.UserInfo, uint64, *errno.Errno) {
	userModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, userModel := range userModels {
		ids = append(ids, userModel.ID)
	}
	info := make([]*apiserverentity.UserInfo, 0)
	wg := sync.WaitGroup{}
	userList := apiserverentity.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*apiserverentity.UserInfo, len(userModels)),
	}
	finished := make(chan bool, 1)
	for _, userModel := range userModels {
		wg.Add(1)
		go func(userModel *apiservermodel.TbUser) {
			defer wg.Done()
			shortId, _ := util.GenShortId()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[userModel.ID] = &apiserverentity.UserInfo{
				ID:          userModel.ID,
				Username:    userModel.Username,
				Mobile:      userModel.Mobile,
				SayHello:    fmt.Sprintf("Hello %s", shortId),
				CreatedTime: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedTime: userModel.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(userModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, userList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a User) List(listParams *service.ListParams) (result []*apiservermodel.TbUser, count int64, err error) {
	qc := apiserverquery.Q.WithContext(a.ctx).TbUser
	if listParams.Options.CustomDBOrder != "" {
		qc = apiserverquery.Q.TbUser.WithContext(a.ctx)
		qc.ReplaceDB(qc.UnderlyingDB().Order(listParams.Options.CustomDBOrder))
	}
	base := qc.Select(listParams.Fields...).Where(listParams.Conditions...).Order(listParams.Orders...)
	offset, limit := util.MysqlPagination(listParams.PS)
	if !listParams.Options.WithoutCount {
		result, count, err = base.FindByPage(offset, limit)
	} else {
		if limit == -1 {
			result, err = base.Find()
		} else {
			result, err = base.Offset(offset).Limit(limit).Find()
		}
	}
	return
}
