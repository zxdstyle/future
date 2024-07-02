package model

import (
	"path/filepath"
	"strings"
	"time"
)

type Object struct {
	ID        string    `json:"id"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Modified  time.Time `json:"modified"`
	Ctime     time.Time `json:"ctime"`
	IsFolder  bool      `json:"is_folder"`
	Thumbnail string    `json:"thumbnail"`
}

func (o Object) IsImage() bool {
	exts := filepath.Ext(o.Path)
	if len(exts) == 0 {
		return false
	}
	ext := strings.ToLower(exts[1:])
	switch ext {
	case "gif",
		"jpg",
		"jpeg",
		"magick",
		"pdf",
		"png",
		"svg",
		"tiff",
		"webp",
		"heif",
		"avif",
		"bmp",
		"jp2k",
		"jxl":
		return true
	default:
		return false
	}
}
