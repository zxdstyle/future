package s3

import (
	"future-admin/internal/logic"
	"future-admin/internal/logic/drivers"
	"future-admin/internal/model"
)

func init() {
	logic.Invoke[*drivers.Logic]().Register(func() model.IDriver {
		return &S3{
			config: model.DriverConfig{
				Slug: "s3",
				Name: "对象存储(S3)",
			},
		}
	})
}
