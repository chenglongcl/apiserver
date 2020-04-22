package ws

import (
	"gopkg.in/olahol/melody.v1"
	"sync"
)

type connMgr struct {
	counter         int64
	sessionWithRoom map[string][]*melody.Session
	rwMutex         *sync.RWMutex
	closeChan       chan *melody.Session
}

var cm *connMgr

func Init() {
	cm = &connMgr{
		counter:         0,
		sessionWithRoom: make(map[string][]*melody.Session),
		rwMutex:         new(sync.RWMutex),
		closeChan:       make(chan *melody.Session, 100),
	}
	go func() {
		for {
			select {
			case s := <-cm.closeChan:
				go func() {
					if !s.IsClosed() {
						s.Close()
					}
					roomID, existRoomID := s.Get("roomID")
					if existRoomID {
						cm.rwMutex.Lock()
						cm.sessionWithRoom[roomID.(string)] =
							removeSession(cm.sessionWithRoom[roomID.(string)], s)
						cm.counter--
						if len(cm.sessionWithRoom[roomID.(string)]) == 0 {
							delete(cm.sessionWithRoom, roomID.(string))
						}
						cm.rwMutex.Unlock()
					}
				}()
			}
		}
	}()
}

func EnterRoom(s *melody.Session) {
	roomID, _ := s.Get("roomID")
	cm.rwMutex.Lock()
	cm.counter++
	cm.sessionWithRoom[roomID.(string)] = append(cm.sessionWithRoom[roomID.(string)], s)
	cm.rwMutex.Unlock()
}

func LeaveRoom(s *melody.Session) {
	cm.closeChan <- s
}

func GetSessionByRoomID(roomID string) []*melody.Session {
	defer cm.rwMutex.RUnlock()
	cm.rwMutex.RLock()
	if _, ok := cm.sessionWithRoom[roomID]; ok {
		return cm.sessionWithRoom[roomID]
	}
	return nil
}

func removeSession(ss []*melody.Session, elem *melody.Session) []*melody.Session {
	var index int
	for k, s := range ss {
		if s == elem {
			index = k
			break
		}
	}
	return append(ss[:index], ss[index+1:]...)
}
