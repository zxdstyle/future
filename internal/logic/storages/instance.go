package storages

import (
	"context"
	"fmt"
	"future-admin/internal/model"
	"future-admin/pkg/log"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"strings"
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
		queue = make(chan *model.Object, 64)
	)

	pool, err := ants.NewPool(runtime.NumCPU())
	if err != nil {
		return err
	}

	go func() {
		if err := i.driver.Walk(ctx, dir, func(obj *model.Object) error {
			queue <- obj
			return nil
		}); err != nil {
			log.Error(err)
		}
		close(queue)
	}()

	for obj := range queue {
		if err := pool.Submit(func() {
			thumbnail, err := i.Thumbnail(obj.Path)
			if err != nil {
				log.Error(err)
			}

			if err = i.driver.Put(ctx, thumbnail); err != nil {
				log.Error(err)
			}

			obj.Thumbnail = fmt.Sprintf("http://127.0.0.1:8081/fs/image?path=%s&storage_id=%d", thumbnail.Name, i.storage.ID)

			if err := handler(obj); err != nil {
				log.Error(err)
			}
		}); err != nil {
			return err
		}
	}

	return nil
}

func (i *Instance) Thumbnail(file string) (*model.FileData, error) {
	content, err := i.driver.Read(file)
	if err != nil {
		return nil, err
	}

	data, err := i.generateThumbnail(content)
	if err != nil {
		return nil, err
	}

	elem := strings.Split(file, "/")
	filename := elem[len(elem)-1]
	eles := strings.Split(filename, ".")

	name := eles[0] + ".avif"

	return &model.FileData{
		Ext:  "avif",
		Name: "thumbnails/" + name,
		Data: data,
		Headers: map[string]string{
			"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, name),
			"Content-Type":        "image/" + eles[1],
		},
	}, nil
}

func (i *Instance) generateThumbnail(content []byte) ([]byte, error) {
	img, err := vips.NewImageFromBuffer(content)
	if err != nil {
		return nil, err
	}
	if err := img.Thumbnail(800, 800, vips.InterestingAttention); err != nil {
		return nil, err
	}
	avif := vips.NewWebpExportParams()
	data, _, err := img.ExportWebp(avif)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (i *Instance) Read(ctx context.Context, file string) ([]byte, error) {
	return i.driver.Read(file)
}
