package handler

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"future-admin/internal/logic"
	"future-admin/internal/logic/storages"
	"future-admin/internal/model"
	"future-admin/pkg/log"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
	"future-admin/pkg/utils/files"
	"future-admin/pkg/utils/images"
	"github.com/gofiber/fiber/v2"
	"io/fs"
	"net/url"
	"os"
	"strings"
)

var Fs = new(fsHandler)

type fsHandler struct {
}

type ListReq struct {
	Path      string `query:"path"`
	StorageId uint64 `query:"storage_id"`
}

func (s *fsHandler) List(ctx context.Context, r *requests.RequestAble) (responses.Response, error) {
	var req ListReq
	if err := r.QueryParser(&req); err != nil {
		return nil, err
	}

	instance, err := logic.Invoke[*storages.Logic]().Instance(ctx, req.StorageId)
	if err != nil {
		return nil, err
	}

	r.Set("Content-Type", "text/event-stream")
	r.Set("Cache-Control", "no-cache")
	r.Set("Connection", "keep-alive")
	r.Set("Transfer-Encoding", "chunked")
	r.Status(fiber.StatusOK).Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		err = instance.Walk(context.Background(), req.Path, func(obj *model.Object) error {
			log.Info(obj.Path)
			data, err := json.Marshal(obj)
			if err != nil {
				log.Error(err)
				return err
			}

			if _, err := fmt.Fprintf(w, "data: %s\n\n", string(data)); err != nil {
				log.Error(err)
				return err
			}

			return w.Flush()
		})
		if err != nil {
			log.Error(err.Error())
		}
		log.Info("ended")
		select {}
	})

	return nil, nil
}

func (s *fsHandler) Detail(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	path := req.Params("path")

	unescape, err := url.QueryUnescape(path)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(unescape)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	desc := s.fsFileInfoToFileDescription(strings.TrimSuffix(unescape, info.Name()), info)

	if !desc.IsDir && files.IsImage(desc.Path) {
		exif, err := files.GetFileExif(desc.Path)
		if err != nil {
			log.Error(err)
		}
		desc.Exif = exif
	}

	return responses.Success(desc), nil
}

func (s *fsHandler) Preview(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	path := req.Params("file")
	file, err := url.QueryUnescape(path)
	if err != nil {
		return nil, err
	}

	img, err := images.Read(file)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	bytes, err := images.Compress(img)
	if err != nil {
		return nil, err
	}
	return responses.Bytes(bytes.Bytes()), nil
}

func (*fsHandler) fsFileInfoToFileDescription(parent string, info fs.FileInfo) model.FileDescription {
	var (
		name      = info.Name()
		createdAt = files.GetFileCreateTime(info)
		updatedAt = files.GetFileUpdateTime(info)
		accessAt  = files.GetFileAccessTime(info)
	)

	return model.FileDescription{
		IsDir:      info.IsDir(),
		Path:       fmt.Sprintf("%s/%s", parent, name),
		Filename:   info.Name(),
		Mode:       info.Mode().String(),
		Size:       info.Size(),
		CreatedAt:  createdAt,
		AccessedAt: accessAt,
		UpdatedAt:  updatedAt,
	}
}

func (s *fsHandler) Image(ctx context.Context, r *requests.RequestAble) (responses.Response, error) {
	var req ListReq
	if err := r.QueryParser(&req); err != nil {
		return nil, err
	}

	instance, err := logic.Invoke[*storages.Logic]().Instance(ctx, req.StorageId)
	if err != nil {
		return nil, err
	}

	data, err := instance.Read(ctx, req.Path)
	if err != nil {
		return nil, err
	}

	return responses.Bytes(data), nil
}
