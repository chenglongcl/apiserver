package msgmgr

import "gopkg.in/olahol/melody.v1"

type MsgMgr interface {
	Broadcast(m *melody.Melody, s *melody.Session)
	BroadcastFilter(m *melody.Melody, s *melody.Session)
}
