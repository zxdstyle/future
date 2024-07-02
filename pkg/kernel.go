package pkg

import (
	"future-admin/pkg/config"
	"future-admin/pkg/database/db"
	"github.com/samber/do"
)

func init() {
	config.Load()

	do.Provide(container, db.NewManager)
}
