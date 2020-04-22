package websocketservice

import (
	"apiserver/ws/msgmgr"
	"gopkg.in/olahol/melody.v1"
)

//Ploy 1房间 2私聊
type Message struct {
	Code    int64
	Content string
	Ploy    int64
}

func (a *Message) Process(m *melody.Melody, s *melody.Session) {
	var msgMgr msgmgr.MsgMgr
	switch a.Ploy {
	case 1:
		msgMgr = &msgmgr.RoomMsgMgr{
			Code:    0,
			Content: a.Content,
		}
		msgMgr.Broadcast(m, s)
	}
}
