package sd

import (
	. "apiserver/handler"
	"apiserver/mq"
	"apiserver/pkg/producermq"
	"apiserver/service/demoservice"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	demoService := &demoservice.Demo{}
	info := demoService.DemoOne()
	SendResponse(c, nil, info)
}

func DemoTwo(c *gin.Context) {
	demoService := &demoservice.Demo{}
	info := demoService.DemoTwo()
	SendResponse(c, nil, info)
}

func DemoThree(c *gin.Context) {
	p := producermq.GetProducer()
	for i := 0; i < 10; i++ {
		go func(i int) {
			msg := mq.NewPublishMsg([]byte(fmt.Sprintf(`{"name":"apiserver-%d"}`, i)))
			_ = p.Publish("exch.unitest", "route.unitest2", msg)
		}(i)
	}
	SendResponse(c, nil, nil)
}
