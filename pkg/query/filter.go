package query

import (
	"github.com/spf13/cast"
)

type (
	Cast func(string) any

	Filter struct {
		Operators map[Operator]struct{}
		Cast      Cast
	}
)

var (
	CastString Cast = func(s string) any { return s }

	CastInt  Cast = func(s string) any { return cast.ToInt(s) }
	CastUint Cast = func(s string) any { return cast.ToUint(s) }

	CastInt32  Cast = func(s string) any { return cast.ToInt32(s) }
	CastUint32 Cast = func(s string) any { return cast.ToUint32(s) }

	CastInt64  Cast = func(s string) any { return cast.ToInt64(s) }
	CastUint64 Cast = func(s string) any { return cast.ToUint64(s) }
)
