package apiserverentity

import "sync"

type UserInfo struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Mobile      string `json:"mobile"`
	SayHello    string `json:"sayHello"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}
