package local

import (
	"context"
	"encoding/json"
	"future-admin/internal/constants/errs"
	"future-admin/internal/model"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/djherbis/times"
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
	rawFiles, err := readDir(dir)
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

func (d *Local) Thumbnail() {
	img, _ := vips.NewImageFromFile()
	img.Thumbnail()
}
