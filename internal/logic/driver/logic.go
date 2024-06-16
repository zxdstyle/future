package driver

import (
	"context"
	"future-admin/internal/model"
	"github.com/samber/do"
)

func New(i *do.Injector) (*Logic, error) {
	return &Logic{}, nil
}

type Logic struct {
}

func (*Logic) List(ctx context.Context) ([]model.DriverInfo, error) {
	return driverInfo.Values(), nil
}
