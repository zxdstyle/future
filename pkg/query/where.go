package query

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Where struct {
	field    string
	operator string
	value    interface{}
}

const (
	whereSlug = "where"
)

func newWhere(segments []string, value any) Builder {
	return &Where{
		field:    segments[1],
		operator: segments[2],
		value:    value,
	}
}

func (w *Where) Build(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Eq{
		Column: w.field,
		Value:  w.value,
	})
}
