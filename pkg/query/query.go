package query

import (
	"context"
	"future-admin/pkg/server/requests"
	"gorm.io/gorm"
	"strings"
)

type Query struct {
	model    any
	builders []Builder
	tx       *gorm.DB
}

func New(model any) *Query {
	return &Query{
		model: model,
	}
}

func (q *Query) WithDB(tx *gorm.DB) *Query {
	q.tx = tx
	return q
}

func (q *Query) ParseFromReq(req *requests.RequestAble) (*Query, error) {
	var (
		queries  = req.Queries()
		builders = make([]Builder, 0)

		filters = make(map[string]Filter, 0)
	)

	if filter, ok := q.model.(WithFilter); ok {
		filters = filter.Filter()
	}

	for key, value := range queries {
		s := strings.Split(key, ".")
		switch s[0] {
		case whereSlug:
			if opt, ok := filters[s[1]]; ok {
				if _, o := opt.Operators[Operator(s[2])]; o {
					builders = append(builders, newWhere(s, opt.Cast(value)))
				}
			}
		case sortSlug:
			builders = append(builders, newSort(s[1], value))
		case pageSlug:
		case pageSizeSlug:
		}
	}

	q.builders = builders
	return q, nil
}

func (q *Query) Do(ctx context.Context) *gorm.DB {
	tx := q.tx.WithContext(ctx)
	for _, builder := range q.builders {
		tx = builder.Build(tx)
	}
	return tx
}
