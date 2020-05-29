package websockets

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Ws(c *gin.Context, m *melody.Melody) {
	//协议升级
	_ = m.HandleRequest(c.Writer, c.Request)
}
