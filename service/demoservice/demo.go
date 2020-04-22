package demoservice

import (
	"apiserver/model"
	"apiserver/pkg/redisgo"
)

type Demo struct {
}

func (a *Demo) DemoOne() (*model.User) {
	user, err := model.GetUser(1)
	if err != nil {
		panic(err)
	}
	_, _ = redisgo.My().HSet("testUsers", "1", user)
	userTwo := model.User{}
	_ = redisgo.My().HGetObject("testUsers", "1", &userTwo)
	return &userTwo
}

func (a *Demo) DemoTwo() int64 {
	_, _ = redisgo.My().IncrBy("testCount", 1)
	testCount, _ := redisgo.My().GetInt64("testCount")
	return testCount
}
