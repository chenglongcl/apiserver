package userservice

import (
	"apiserver/model"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/util"
	"fmt"
	"sync"
)

type User struct {
	ID       uint64
	Username string
	Password string
	Mobile   string
}

func (a *User) Add() (uint64, *errno.Errno) {
	if userExist, _ := model.CheckUserByUsername(a.Username); userExist {
		return 0, errno.ErrUserExist
	}
	password, _ := auth.Encrypt(a.Password)
	data := map[string]interface{}{
		"username": a.Username,
		"password": password,
		"mobile":   a.Mobile,
	}
	id, err := model.AddUser(data)
	if err != nil {
		return 0, errno.ErrDatabase
	}
	return id, nil
}

func (a *User) Get() (*model.User, *errno.Errno) {
	user, err := model.GetUser(a.ID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return user, nil
}

func (a *User) Edit() *errno.Errno {
	if userExist, _ := model.CheckUserByID(a.ID); !userExist {
		return errno.ErrNotUserExist
	}
	var password string
	if a.Password != "" {
		password, _ = auth.Encrypt(a.Password)
	}
	data := map[string]interface{}{
		"id":       a.ID,
		"password": password,
		"mobile":   a.Mobile,
	}
	err := model.EditUser(data)
	if err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *User) Delete() *errno.Errno {
	if err := model.DeleteUser(a.ID); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *User) GetUserList(ps util.PageSetting) ([]*model.UserInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{})
	if a.Username != "" {
		w["username like"] = "%" + a.Username + "%"
	}
	users, count, err := model.GetUserList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	info := make([]*model.UserInfo, 0)
	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	finished := make(chan bool, 1)
	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()
			shortId, _ := util.GenShortId()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.ID] = &model.UserInfo{
				ID:          u.ID,
				Username:    u.Username,
				Mobile:      u.Mobile,
				SayHello:    fmt.Sprintf("Hello %s", shortId),
				CreatedTime: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedTime: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, userList.IdMap[id])
	}
	return info, count, nil
}
