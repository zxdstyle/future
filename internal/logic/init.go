package logic

import (
	"future-admin/internal/logic/album"
	"github.com/samber/do"
)

func init() {
	do.Provide(nil, album.New)
}

func Invoke[M any]() M {
	return do.MustInvoke[M](nil)
}
