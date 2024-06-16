package logic

import (
	"future-admin/internal/logic/album"
	"future-admin/internal/logic/driver"
	"future-admin/internal/logic/storage"
	"github.com/samber/do"
)

func init() {
	do.Provide(nil, album.New)
	do.Provide(nil, driver.New)
	do.Provide(nil, storage.New)
}

func Invoke[M any]() M {
	return do.MustInvoke[M](nil)
}
