package logic

import (
	"future-admin/internal/dao"
	"future-admin/internal/logic/album"
	"future-admin/internal/logic/drivers"
	"future-admin/internal/logic/storages"
	"future-admin/pkg"
	"github.com/samber/do"
)

func init() {
	dao.SetDefault(pkg.DBManager().DB())

	do.Provide(nil, album.New)
	do.Provide(nil, drivers.New)
	do.Provide(nil, storages.New)
}

func Invoke[M any]() M {
	return do.MustInvoke[M](nil)
}
