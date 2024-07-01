package storages

import (
	"context"
	"fmt"
	"future-admin/internal/model"
	"github.com/davidbyttow/govips/v2/vips"
	"golang.org/x/sync/errgroup"
	"strings"
)

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

func (i *Instance) Walk(ctx context.Context, dir string, handler func(obj *model.Object) error) error {
	var (
		queue = make(chan *model.Object, 64)
		group = errgroup.Group{}
	)

	group.Go(func() error {
		if err := i.driver.Walk(ctx, dir, func(obj *model.Object) error {
			queue <- obj
			return nil
		}); err != nil {
			return err
		}
		close(queue)
		return nil
	})

	for obj := range queue {
		group.Go(func() error {
			return handler(obj)
		})
	}

	return group.Wait()
}

func (i *Instance) Thumbnail(file string) (*model.FileData, error) {
	content, err := i.driver.Read(file)
	if err != nil {
		return nil, err
	}
	img, err := vips.NewImageFromBuffer(content)
	if err != nil {
		return nil, err
	}
	if err := img.Thumbnail(800, 800, vips.InterestingAttention); err != nil {
		return nil, err
	}
	avif := vips.NewAvifExportParams()
	avif.Quality = 80
	avif.Effort = 8

	data, meta, err := img.ExportAvif(avif)
	if err != nil {
		return nil, err
	}

	elem := strings.Split(file, "/")
	eles := strings.Split(elem[len(elem)-1], ".")
	filename := eles[0] + meta.Format.FileExt()

	return &model.FileData{
		Ext:  "avif",
		Name: filename,
		Data: data,
		Headers: map[string]string{
			"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, filename),
			"Content-Type":        "image/" + meta.Format.FileExt(),
		},
	}, nil
}
