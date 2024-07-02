package drivers

import (
	"future-admin/drivers/local"
	"future-admin/drivers/s3"
	"future-admin/internal/logic"
	"future-admin/internal/logic/drivers"
)

func init() {
	logic.Invoke[*drivers.Logic]().Register(s3.Register)
	logic.Invoke[*drivers.Logic]().Register(local.Register)
}
