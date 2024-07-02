package model

import "github.com/golang-module/carbon/v2"

type FileDescription struct {
	IsDir      bool          `json:"is_dir"`
	Path       string        `json:"path"`
	Filename   string        `json:"filename"`
	Mode       string        `json:"mode"`
	Size       int64         `json:"size"`
	CreatedAt  carbon.Carbon `json:"created_at"`
	AccessedAt carbon.Carbon `json:"accessed_at"`
	UpdatedAt  carbon.Carbon `json:"updated_at"`

	Exif []ExifItem `json:"exif,omitempty"`
}
