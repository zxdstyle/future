package model

import "context"

type (
	IDriver interface {
		List(dir string) ([]*Object, error)
		Config() DriverConfig
		Init(addition string) error
		GetAddition() any
		Read(file string) ([]byte, error)
		Walk(ctx context.Context, dir string, handler func(obj *Object) error) error
		Put(ctx context.Context, file *FileData) error
	}

	DriverConfig struct {
		Name     string
		Slug     string
		PageSize int
	}

	DriverInfo struct {
		Slug      string               `json:"slug"`
		Name      string               `json:"name"`
		Additions []DriverAdditionItem `json:"additions"`
	}

	DriverAdditionItem struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Default  string `json:"default"`
		Options  string `json:"options"`
		Required bool   `json:"required"`
		Help     string `json:"help"`
	}
)
