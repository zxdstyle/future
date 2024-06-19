package storages

import "future-admin/internal/model"

type Instance struct {
	driver model.IDriver
}

func newInstance(storage *model.Storage, driver model.IDriver) (*Instance, error) {
	if err := driver.Init(storage.Addition); err != nil {
		return nil, err
	}
	return &Instance{
		driver,
	}, nil
}

func (i *Instance) GetList() error {
	return i.driver.List()
}
