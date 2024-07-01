package local

import (
	"future-admin/internal/model"
)

func Register() model.IDriver {
	return &Local{
		config: model.DriverConfig{
			Slug: "local",
			Name: "本地存储",
		},
	}
}
