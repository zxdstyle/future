package query

import "gorm.io/gorm"

type Builder interface {
	Build(tx *gorm.DB) *gorm.DB
}
