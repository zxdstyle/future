package query

type Operator string

var (
	Eq      Operator = "eq"
	Neq     Operator = "neq"
	Gt      Operator = "gt"
	Gte     Operator = "gte"
	Lt      Operator = "lt"
	Lte     Operator = "lte"
	In      Operator = "in"
	Like    Operator = "like"
	Between Operator = "between"
)

func WithOperator(ops ...Operator) map[Operator]struct{} {
	result := make(map[Operator]struct{})
	for _, op := range ops {
		result[op] = struct{}{}
	}
	return result
}
