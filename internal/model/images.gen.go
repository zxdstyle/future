// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/golang-module/carbon/v2"
)

const TableNameImage = "images"

// Image mapped from table <images>
type Image struct {
	ID        uint64         `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Path      string         `gorm:"column:path;type:varchar(255);not null" json:"path"`
	Ext       string         `gorm:"column:ext;type:varchar(32);not null" json:"ext"`
	Size      int32          `gorm:"column:size;type:int;not null" json:"size"`
	CreatedAt *carbon.Carbon `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt *carbon.Carbon `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName Image's table name
func (*Image) TableName() string {
	return TableNameImage
}
