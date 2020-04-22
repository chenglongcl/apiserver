package websockets

import (
	"apiserver/ws"
	"gopkg.in/olahol/melody.v1"
)

func Disconnect(s *melody.Session) {
	ws.LeaveRoom(s)
}
