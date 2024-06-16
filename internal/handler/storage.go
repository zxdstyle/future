package handler

import (
	"future-admin/internal/logic"
	"future-admin/internal/logic/storage"
	"future-admin/internal/model"
	"future-admin/pkg/base"
)

var Storage = &storageHandler{
	Handler: base.NewHandler[model.Storage](
		logic.Invoke[*storage.Logic]().Logic,
	),
}

type storageHandler struct {
	*base.Handler[model.Storage]
}
