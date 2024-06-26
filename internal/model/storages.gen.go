// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/golang-module/carbon/v2"
)

const TableNameStorage = "storages"

// Storage mapped from table <storages>
type Storage struct {
	ID        uint64         `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      *string        `gorm:"column:name;type:varchar(255)" json:"name"`
	Driver    string         `gorm:"column:driver;type:varchar(64);not null" json:"driver"`
	Addition  string         `gorm:"column:addition;type:text;not null" json:"addition"`
	Remark    *string        `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreatedAt *carbon.Carbon `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt *carbon.Carbon `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName Storage's table name
func (*Storage) TableName() string {
	return TableNameStorage
}
