package local

import (
	"context"
	"encoding/json"
	"errors"
	"future-admin/internal/constants/errs"
	"future-admin/internal/model"
	"github.com/djherbis/times"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Local struct {
	Addition
	config model.DriverConfig
}

func (d *Local) Init(addition string) error {
	if err := json.Unmarshal([]byte(addition), &d.Addition); err != nil {
		return err
	}

	return nil
}

func (d *Local) Config() model.DriverConfig {
	return d.config
}

func (d *Local) GetAddition() any {
	return d.Addition
}

func (d *Local) List(dir string) ([]*model.Object, error) {
	rawFiles, err := readDir(path.Join(d.Path, dir))
	if err != nil {
		return nil, err
	}
	var files []*model.Object
	for _, f := range rawFiles {
		if !d.ShowHidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		file := d.FileInfoToObj(f, dir)
		files = append(files, file)
	}
	return files, nil
}

func (d *Local) FileInfoToObj(f fs.FileInfo, fullPath string) *model.Object {
	//thumb := ""
	//if d.Thumbnail {
	//	typeName := utils.GetFileType(f.Name())
	//	if typeName == conf.IMAGE || typeName == conf.VIDEO {
	//		thumb = common.GetApiUrl(nil) + stdpath.Join("/d", reqPath, f.Name())
	//		thumb = utils.EncodePath(thumb, true)
	//		thumb += "?type=thumb&sign=" + sign.Sign(stdpath.Join(reqPath, f.Name()))
	//	}
	//}

	isFolder := f.IsDir() || isSymlinkDir(f, fullPath)
	var size int64
	if !isFolder {
		size = f.Size()
	}
	var ctime time.Time
	t, err := times.Stat(path.Join(fullPath, f.Name()))
	if err == nil {
		if t.HasBirthTime() {
			ctime = t.BirthTime()
		}
	}

	//file := model.ObjThumb{
	//	Object: model.Object{
	//		Path:     filepath.Join(fullPath, f.Name()),
	//		Name:     f.Name(),
	//		Modified: f.ModTime(),
	//		Size:     size,
	//		IsFolder: isFolder,
	//		Ctime:    ctime,
	//	},
	//	Thumbnail: model.Thumbnail{
	//		Thumbnail: thumb,
	//	},
	//}

	return &model.Object{
		ID:       "",
		Path:     filepath.Join(fullPath, f.Name()),
		Name:     f.Name(),
		Size:     size,
		Modified: f.ModTime(),
		Ctime:    ctime,
		IsFolder: isFolder,
	}
}

func (d *Local) Get(ctx context.Context, path string) (*model.Object, error) {
	f, err := os.Stat(path)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the file") {
			return nil, errs.ErrObjectNotFound
		}
		return nil, err
	}
	isFolder := f.IsDir() || isSymlinkDir(f, path)
	size := f.Size()
	if isFolder {
		size = 0
	}
	var ctime time.Time
	t, err := times.Stat(path)
	if err == nil {
		if t.HasBirthTime() {
			ctime = t.BirthTime()
		}
	}
	file := model.Object{
		Path:     path,
		Name:     f.Name(),
		Modified: f.ModTime(),
		Ctime:    ctime,
		Size:     size,
		IsFolder: isFolder,
	}
	return &file, nil
}

func (d *Local) Walk(ctx context.Context, dir string, handler func(obj *model.Object) error) error {
	fullPath := path.Join(d.Path, dir)

	fs, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer fs.Close()

	info, err := fs.Stat()
	if err != nil {
		return err
	}
	if !info.IsDir() {
		if !d.ShowHidden && strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		obj := d.FileInfoToObj(info, dir)
		return handler(obj)
	}

	for {
		files, err := fs.ReadDir(30)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				return err
			}

			if !d.ShowHidden && strings.HasPrefix(info.Name(), ".") {
				continue
			}

			obj := d.FileInfoToObj(info, dir)
			if err := handler(obj); err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Local) Read(file string) ([]byte, error) {
	return os.ReadFile(path.Join(d.Path, file))
}

func (d *Local) Put(ctx context.Context, file *model.FileData) error {
	fullPath := path.Join(d.Path, file.Name)
	out, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
		if errors.Is(err, context.Canceled) {
			_ = os.Remove(fullPath)
		}
	}()

	if _, err = out.Write(file.Data); err != nil {
		return err
	}

	return nil
}
