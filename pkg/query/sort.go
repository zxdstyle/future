package query

import (
	"fmt"
	"gorm.io/gorm"
)

const sortSlug = "sort"

type Sort struct {
	field string
	order string
}

func newSort(field, order string) Builder {
	return &Sort{
		field: field,
		order: order,
	}
}

func (s *Sort) Build(tx *gorm.DB) *gorm.DB {
	if s.order == "asc" {
		return tx.Order(fmt.Sprintf("`%s` ASC", s.field))
	}
	return tx.Order(fmt.Sprintf("`%s` DESC", s.field))
}
