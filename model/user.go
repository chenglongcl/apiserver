package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type User struct {
	BaseModel
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Mobile   string `gorm:"column:mobile"`
}

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

func (u *User) TableName() string {
	return viper.GetString("db.prefix") + "user"
}

func CheckUserByUsername(username string) (bool, error) {
	var user User
	err := SelectDB("self").Select("id").Where("username = ?", username).
		Unscoped().First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckUserByID(id uint64) (bool, error) {
	var user User
	err := SelectDB("self").Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddUser(data map[string]interface{}) (uint64, error) {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Mobile:   data["mobile"].(string),
	}
	if err := SelectDB("self").Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GetUser(id uint64) (*User, error) {
	var user User
	err := SelectDB("self").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	d := SelectDB("self").Where("username = ?", username).First(&user)
	return &user, d.Error
}

func EditUser(data map[string]interface{}) error {
	var user User
	if data["password"].(string) == "" {
		SelectDB("self").Model(&user).Omit("password").Update(data)
	} else {
		SelectDB("self").Model(&user).Update(data)
	}
	return nil
}

func DeleteUser(id uint64) error {
	var user User
	return SelectDB("self").Where("id = ?", id).Delete(&user).Error
}

func GetUserList(w map[string]interface{}, fields string, offset, limit uint64, order string) ([]*User, uint64, error) {
	users := make([]*User, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	modelDB := SelectDB("self").Model(&User{})
	if err := modelDB.Where(where, values...).Count(&count).Error; err != nil {
		return users, count, err
	}
	if offset != 0 {
		modelDB = modelDB.Offset(offset)
	}
	if limit != 0 {
		modelDB = modelDB.Limit(limit)
	}
	if err := modelDB.Select(fields).Where(where, values...).Order(order).Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}
