package handler

import (
	"context"
	"errors"
	"fmt"
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

var Storage = new(storage)

type storage struct {
}

func (s *storage) List(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	dir := req.Query("query[parent_dir]")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	desc := make([]model.FileDescription, 0)
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		var (
			name = info.Name()
		)

		if strings.HasPrefix(name, ".") {
			continue
		}

		desc = append(desc, s.fsFileInfoToFileDescription(dir, info))
	}

	return responses.Success(files.SortFiles(desc)), nil
}

func (s *storage) Detail(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
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

func (s *storage) Preview(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
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
	return responses.Bytes(bytes), nil
}

func (*storage) fsFileInfoToFileDescription(parent string, info fs.FileInfo) model.FileDescription {
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
