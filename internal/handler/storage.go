package handler

import (
	"future-admin/internal/logic"
	"future-admin/internal/logic/storages"
	"future-admin/internal/model"
	"future-admin/pkg/base"
)

var Storage = &storageHandler{
	Handler: base.NewHandler[model.Storage](
		logic.Invoke[*storages.Logic]().Logic,
	),
}

type storageHandler struct {
	*base.Handler[model.Storage]
}
