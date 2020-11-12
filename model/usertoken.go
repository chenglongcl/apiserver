package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"time"
)

type UserToken struct {
	UserID      uint64     `gorm:"primary_key;column:user_id;type:int(11);not null" mapstructure:"user_id"` //
	Token       string     `gorm:"unique;column:token;type:varchar(100);not null"`                          //	token
	ExpireTime  time.Time  `gorm:"column:expire_time;type:datetime" mapstructure:"expire_time"`             //	过期时间
	RefreshTime time.Time  `gorm:"column:refresh_time;type:datetime" mapstructure:"refresh_time"`           //	更新时间
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp"`                                        //	创建时间
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp"`                                        //
	DeletedAt   *time.Time `gorm:"index;column:deleted_at;type:timestamp"`                                  //
}

func (u *UserToken) TableName() string {
	return viper.GetString("db.prefix") + "user_token"
}

func AddUserToken(data map[string]interface{}) error {
	var userToken UserToken
	if err := mapstructure.Decode(data, &userToken); err != nil {
		return err
	}
	if err := SelectDB("self").Create(&userToken).Error; err != nil {
		return err
	}
	return nil
}

func GetUserToken(userID uint64) (*UserToken, error) {
	var userToken UserToken
	err := SelectDB("self").Where("user_id = ?", userID).First(&userToken).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &userToken, nil
}

func EditUserToken(data map[string]interface{}) error {
	var userToken UserToken
	return SelectDB("self").Model(&userToken).Updates(data).Error
}
