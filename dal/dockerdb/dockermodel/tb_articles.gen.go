// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dockermodel

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbArticle = "tb_articles"

// TbArticle mapped from table <tb_articles>
type TbArticle struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	UID       uint64         `gorm:"column:uid;type:int(11);not null" json:"uid"`
	CateID    uint64         `gorm:"column:cate_id;type:int(11);not null" json:"cateID"`
	Title     string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	Images    string         `gorm:"column:images;type:text;not null" json:"images"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
	Editor    string         `gorm:"-"`
}

// TableName TbArticle's table name
func (*TbArticle) TableName() string {
	return TableNameTbArticle
}
