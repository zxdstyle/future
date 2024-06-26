package logic

import (
	"future-admin/internal/dao"
	"future-admin/internal/logic/albums"
	"future-admin/internal/logic/drivers"
	"future-admin/internal/logic/images"
	"future-admin/internal/logic/storages"
	"future-admin/pkg"
	"github.com/samber/do"
)

func init() {
	dao.SetDefault(pkg.DBManager().DB())

	do.Provide(nil, albums.New)
	do.Provide(nil, images.New)
	do.Provide(nil, drivers.New)
	do.Provide(nil, storages.New)
}

func Invoke[M any]() M {
	return do.MustInvoke[M](nil)
}
