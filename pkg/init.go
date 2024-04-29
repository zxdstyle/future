package pkg

import (
	"future-admin/pkg/database/db"
	"github.com/samber/do"
)

var (
	container = do.New()
)

func App[T any]() T {
	return do.MustInvoke[T](container)
}

func DBManager() *db.Manager {
	return do.MustInvoke[*db.Manager](container)
}
