package msgmgr

import (
	"apiserver/ws"
	"github.com/json-iterator/go"
	"gopkg.in/olahol/melody.v1"
)

type RoomMsgMgr struct {
	Code    int64  `json:"code"`
	Content string `json:"content"`
}

func (a *RoomMsgMgr) Broadcast(m *melody.Melody, s *melody.Session) {
	if roomID, exist := s.Get("roomID"); exist {
		if sessions := ws.GetSessionByRoomID(roomID.(string)); m != nil && len(sessions) > 0 {
			data, _ := jsoniter.Marshal(a)
			m.BroadcastMultiple(data, sessions)
		}
	}
}

func (a *RoomMsgMgr) BroadcastFilter(m *melody.Melody, s *melody.Session) {

}
