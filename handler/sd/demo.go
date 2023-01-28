package sd

import (
	"apiserver/handler"
	"apiserver/mq"
	"apiserver/pkg/producermq"
	"apiserver/service/demoservice"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	demoService := demoservice.NewDemoService(c)
	info, errNo := demoService.DemoOne()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}

func DemoTwo(c *gin.Context) {
	demoService := demoservice.NewDemoService(c)
	info := demoService.DemoTwo()
	handler.SendResponse(c, nil, info)
}

func DemoThree(c *gin.Context) {
	p := producermq.GetProducer()
	for i := 0; i < 10; i++ {
		go func(i int) {
			msg := mq.NewPublishMsg([]byte(fmt.Sprintf(`{"name":"apiserver-%d"}`, i)))
			_ = p.Publish("exch.unitest", "route.unitest2", msg)
		}(i)
	}
	handler.SendResponse(c, nil, nil)
}
