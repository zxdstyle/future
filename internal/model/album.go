package model

import (
	"future-admin/pkg/query"
)

func (*Album) Filter() map[string]query.Filter {
	return map[string]query.Filter{
		"id": {
			Operators: query.WithOperator(query.Eq, query.Gt, query.Lt, query.Gte, query.Lte, query.Between),
			Cast:      query.CastInt64},
	}
}
