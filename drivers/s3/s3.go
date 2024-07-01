package s3

import (
	"future-admin/internal/model"
)

func Register() model.IDriver {
	return &S3{
		config: model.DriverConfig{
			Slug: "s3",
			Name: "对象存储(S3)",
		},
	}
}
