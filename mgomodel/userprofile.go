package mgomodel

import (
	"github.com/globalsign/mgo/bson"
)

type UserProfile struct {
	PublicFields `bson:",inline"`
	UserID       uint64 `bson:"user_id" mapstructure:"user_id"`
	RealName     string `bson:"real_name" mapstructure:"real_name"`
	Sex          int64  `bson:"address"`
}

func UpsertUserProfile(data map[string]interface{}) error {
	return GetSession("self").DB("apiserver").Collection("user_profile").
		Upsert(bson.M{"user_id": data["user_id"]}, data)
}
