// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/golang-module/carbon/v2"
)

const TableNameAlbum = "albums"

// Album mapped from table <albums>
type Album struct {
	ID          uint64         `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Title       string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	SubTitle    string         `gorm:"column:sub_title;type:varchar(255);not null" json:"sub_title"`
	Description *string        `gorm:"column:description;type:text" json:"description"`
	CreatedAt   *carbon.Carbon `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   *carbon.Carbon `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName Album's table name
func (*Album) TableName() string {
	return TableNameAlbum
}