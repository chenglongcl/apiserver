package websockets

import (
	"apiserver/ws"
	"gopkg.in/olahol/melody.v1"
)

func Connect(s *melody.Session) {
	roomID := s.Request.FormValue("roomID")
	uid := s.Request.FormValue("uid")
	//TODO AUTH
	s.Set("uid", uid)
	s.Set("roomID", roomID)
	s.Set("isLogin", true)
	ws.EnterRoom(s)
}
