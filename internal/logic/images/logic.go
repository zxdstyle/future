package images

import (
	"future-admin/internal/model"
	"future-admin/pkg"
	"future-admin/pkg/base"
	"github.com/samber/do"
)

func New(i *do.Injector) (*Logic, error) {
	return &Logic{
		Logic: base.NewLogic[model.Image](pkg.DBManager().DB()),
	}, nil
}

type Logic struct {
	*base.Logic[model.Image]
}
