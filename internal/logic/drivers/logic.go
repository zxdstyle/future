package drivers

import (
	"context"
	"future-admin/internal/model"
	"future-admin/pkg/utils/co"
	"github.com/samber/do"
)

func New(i *do.Injector) (*Logic, error) {
	return &Logic{
		drivers:    co.NewMap[string, model.IDriver](),
		driverInfo: co.NewMap[string, model.DriverInfo](),
	}, nil
}

type Logic struct {
	drivers    *co.Map[string, model.IDriver]
	driverInfo *co.Map[string, model.DriverInfo]
}

func (*Logic) List(ctx context.Context) ([]model.DriverInfo, error) {
	return driverInfo.Values(), nil
}
