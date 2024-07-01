package model

import "time"

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
