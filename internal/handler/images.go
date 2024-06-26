package handler

import (
	"context"
	"fmt"
	"future-admin/internal/dao"
	"future-admin/internal/model"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
)

var Image = &imageHandler{}

type imageHandler struct {
}

func (h *imageHandler) Create(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	form, err := req.MultipartForm()
	if err != nil {
		return nil, err
	}

	var (
		files = form.File["file"]
		imgs  = make([]*model.Image, 0)
	)

	for _, file := range files {
		path := fmt.Sprintf("./images/%s", file.Filename)
		if err := req.SaveFile(file, path); err != nil {
			return nil, err
		}

		imgs = append(imgs, &model.Image{
			Path: path,
			Ext:  file.Filename,
			Size: int32(file.Size),
		})
	}

	if err := dao.Image.Create(imgs...); err != nil {
		return nil, err
	}

	return responses.Success("images"), nil
}
