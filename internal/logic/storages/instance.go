package storages

import (
	"context"
	"fmt"
	"future-admin/internal/model"
	"future-admin/pkg/log"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"sync"
)

type Instance struct {
	driver  model.IDriver
	storage *model.Storage
}

func newInstance(storage *model.Storage, driver model.IDriver) (*Instance, error) {
	if err := driver.Init(storage.Addition); err != nil {
		return nil, err
	}
	return &Instance{
		driver, storage,
	}, nil
}

func (i *Instance) Walk(ctx context.Context, dir string, handler func(obj *model.Object) error) error {
	var (
		queue  = make(chan *model.Object, 64)
		result = make(chan *model.Object, 32)
	)

	pool, err := ants.NewPool(runtime.NumCPU())
	if err != nil {
		return err
	}
	defer pool.Release()

	go func() {
		if err := i.driver.Walk(ctx, dir, func(obj *model.Object) error {
			queue <- obj
			return nil
		}); err != nil {
			log.Error(err)
		}
		close(queue)
	}()

	go func() {
		wg := sync.WaitGroup{}
		for obj := range queue {
			wg.Add(1)
			if err := pool.Submit(func() {
				defer wg.Done()
				if obj.IsFolder || !obj.IsImage() {
					result <- obj
					return
				}

				thumbnail, err := i.Thumbnail(obj.Path)
				if err != nil {
					log.Error(err)
				}

				if err = i.driver.Put(ctx, thumbnail); err != nil {
					log.Error(err)
				}

				obj.Thumbnail = fmt.Sprintf("http://127.0.0.1:8081/fs/image?path=%s&storage_id=%d", thumbnail.Name, i.storage.ID)

				result <- obj
			}); err != nil {
				log.Error(err)
			}
		}
		wg.Wait()
		close(result)
	}()

	for obj := range result {
		if err := handler(obj); err != nil {
			log.Error(err)
		}
	}

	return nil
}

func (i *Instance) Read(ctx context.Context, file string) ([]byte, error) {
	return i.driver.Read(file)
}
