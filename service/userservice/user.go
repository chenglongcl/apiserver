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
		user *apiservermodel.TbUser
		err  error
	)
	if count, _ := qc.TbUser.Where(q.TbUser.Username.Eq(a.Username)).Unscoped().Count(); count > 0 {
		return user, errno.ErrUserExist
	}
	password, _ := auth.Encrypt(a.Password)
	user = &apiservermodel.TbUser{
		Username: a.Username,
		Password: password,
		Mobile:   a.Mobile,
		Status:   true,
	}
	err = qc.TbUser.Create(user)
	return user, gormx.HandleError(err)
}

// Get
// @Description:
// @receiver a
// @param fields
// @param conditions
// @return *apiservermodel.TbUser
// @return *errno.Errno

func (a User) Get(fields []field.Expr, conditions []gen.Condition) (*apiservermodel.TbUser, *errno.Errno) {
	user, err := apiserverquery.Q.WithContext(a.ctx).TbUser.Select(fields...).Where(conditions...).Take()
	return user, gormx.HandleError(err)
}

// GetByID
// @Description:
// @receiver a
// @return *apiservermodel.TbUser
// @return *errno.Errno

func (a User) GetByID() (*apiservermodel.TbUser, *errno.Errno) {
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	user, err := qc.TbUser.Where(q.TbUser.ID.Eq(a.ID)).Take()
	return user, gormx.HandleError(err)
}

// GetByUsername
// @Description:
// @receiver a
// @return *apiservermodel.TbUser
// @return *errno.Errno

func (a User) GetByUsername() (*apiservermodel.TbUser, *errno.Errno) {
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	user, err := qc.TbUser.Where(q.TbUser.Username.Eq(a.Username)).Take()
	return user, gormx.HandleError(err)
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
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	_, err := qc.TbUser.Where(q.TbUser.ID.Eq(a.ID)).Delete()
	return gormx.HandleError(err)
}

// GetUserList
// @Description:
// @receiver a
// @param ps
// @param fields
// @param conditions
// @param orders
// @return []*apiserverentity.UserInfo
// @return uint64
// @return *errno.Errno

func (a User) GetUserList(ps util.PageSetting, fields []field.Expr, conditions []gen.Condition, orders []field.Expr) ([]*apiserverentity.UserInfo, uint64, *errno.Errno) {
	qc := apiserverquery.Q.WithContext(a.ctx)
	users, count, err := qc.TbUser.Select(fields...).Where(conditions...).Order(orders...).FindByPage(int(ps.Offset), int(ps.Limit))
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	info := make([]*apiserverentity.UserInfo, 0)
	wg := sync.WaitGroup{}
	userList := apiserverentity.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*apiserverentity.UserInfo, len(users)),
	}
	finished := make(chan bool, 1)
	for _, u := range users {
		wg.Add(1)
		go func(user *apiservermodel.TbUser) {
			defer wg.Done()
			shortId, _ := util.GenShortId()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[user.ID] = &apiserverentity.UserInfo{
				ID:          user.ID,
				Username:    user.Username,
				Mobile:      user.Mobile,
				SayHello:    fmt.Sprintf("Hello %s", shortId),
				CreatedTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedTime: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
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
