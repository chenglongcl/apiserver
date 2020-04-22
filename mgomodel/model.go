package mgomodel

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type PublicFields struct {
	ID        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// SetFieldsValue 设置公共字段值，在插入数据时使用
func (a *PublicFields) SetFieldsValue() {
	now := time.Now()
	if !a.ID.Valid() {
		a.ID = bson.NewObjectId()
		if a.CreatedAt.IsZero() {
			a.CreatedAt = now
		}
	}
	if a.UpdatedAt.IsZero() {
		a.UpdatedAt = now
	}
}
