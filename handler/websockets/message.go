package websockets

import (
	"apiserver/service/websocketservice"
	"apiserver/util"
	"github.com/json-iterator/go"
	"gopkg.in/olahol/melody.v1"
)

func Message(m *melody.Melody, s *melody.Session, msg []byte) {
	var receive Receive
	_ = jsoniter.Unmarshal(msg, &receive)
	if err := util.Validate(&receive); err != nil {
		return
	}
	messageService := websocketservice.Message{
		Content: receive.Content,
		Ploy:    receive.Ploy,
	}
	messageService.Process(m, s)
}
