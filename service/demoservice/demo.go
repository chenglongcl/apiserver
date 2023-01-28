package demoservice

import (
	"apiserver/dal/apiserverdb/apiservermodel"
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/pkg/errno"
	"apiserver/pkg/gormx"
	"apiserver/pkg/redisgo"
	"apiserver/service"
	"github.com/gin-gonic/gin"
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
func (a *Demo) DemoOne() (*apiservermodel.TbUser, *errno.Errno) {
	q := apiserverquery.Q
	qc := apiserverquery.Q.WithContext(a.ctx)
	user, err := qc.TbUser.Where(q.TbUser.ID.Eq(1)).Take()
	if errNo := gormx.HandleError(err); errNo != nil {
		return user, errNo
	}
	_, _ = redisgo.My().HSet("testUsers", "1", user)
	userTwo := apiservermodel.TbUser{}
	_ = redisgo.My().HGetObject("testUsers", "1", &userTwo)
	return &userTwo, nil
}

func (a *Demo) DemoTwo() int64 {
	_, _ = redisgo.My().IncrBy("testCount", 1)
	testCount, _ := redisgo.My().GetInt64("testCount")
	return testCount
}
