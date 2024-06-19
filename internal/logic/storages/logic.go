package storages

import (
	"context"
	"future-admin/internal/dao"
	"future-admin/internal/logic/drivers"
	"future-admin/internal/model"
	"future-admin/pkg"
	"future-admin/pkg/base"
	"future-admin/pkg/utils/co"
	"github.com/samber/do"
)

func New(i *do.Injector) (*Logic, error) {
	return &Logic{
		Logic:       base.NewLogic[model.Storage](pkg.DBManager().DB()),
		instances:   co.NewMap[uint64, *Instance](),
		driverLogic: do.MustInvoke[*drivers.Logic](i),
	}, nil
}

type Logic struct {
	*base.Logic[model.Storage]
	instances   *co.Map[uint64, *Instance]
	driverLogic *drivers.Logic
}

func (l *Logic) Instance(ctx context.Context, id uint64) (*Instance, error) {
	instance, ok := l.instances.Get(id)
	if ok {
		return instance, nil
	}

	instance, err := l.createStorageInstance(ctx, id)
	if err != nil {
		return nil, err
	}

	return instance, nil

}

func (l *Logic) createStorageInstance(ctx context.Context, id uint64) (*Instance, error) {
	storage, err := dao.Storage.Where(dao.Storage.ID.Eq(id)).Take()
	if err != nil {
		return nil, err
	}

	driver, err := l.driverLogic.GetDriver(storage.Driver)
	if err != nil {
		return nil, err
	}

	return newInstance(storage, driver)
}
