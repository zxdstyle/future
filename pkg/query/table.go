package query

type (
	TableModel interface {
		TableName() string
	}

	WithFilter interface {
		Filter() map[string]Filter
	}
)
