package demoservice

import (
	"apiserver/dal/apiserverdb/apiserverentity"
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/pkg/errno"
	"apiserver/pkg/redisgo"
	"apiserver/service"
	"apiserver/service/userservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type Demo struct {
	serviceOptions *service.Options
	ctx            *gin.Context
}

func NewDemoService(ctx *gin.Context, opts ...service.Option) *Demo {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &Demo{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a *Demo) DemoOne() (*apiserverentity.UserInfo, *errno.Errno) {
	userService := userservice.NewUserService(a.ctx)
	user, errNo := userService.Get([]field.Expr{
		apiserverquery.Q.TbUser.ALL,
	}, []gen.Condition{
		apiserverquery.Q.TbUser.ID.Eq(1),
	})
	if errNo != nil {
		return nil, errNo
	}
	if user == nil || user.ID == 0 {
		return nil, errno.ErrUserNotFound
	}
	_, _ = redisgo.My().HSet("testUsers", "1", apiserverentity.UserInfo{
		ID:          user.ID,
		Username:    user.Username,
		Mobile:      user.Mobile,
		SayHello:    "hello",
		CreatedTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedTime: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
	userTwo := apiserverentity.UserInfo{}
	_ = redisgo.My().HGetObject("testUsers", "1", &userTwo)
	return &userTwo, nil
}

func (a *Demo) DemoTwo() int64 {
	_, _ = redisgo.My().IncrBy("testCount", 1)
	testCount, _ := redisgo.My().GetInt64("testCount")
	return testCount
}
