package handler

import (
	"future-admin/internal/logic"
	"future-admin/internal/logic/album"
	"future-admin/internal/model"
	"future-admin/pkg/base"
)

var Album = &albumHandler{
	Handler: base.NewHandler[model.Album](
		logic.Invoke[*album.Logic]().Logic,
	),
}

type albumHandler struct {
	*base.Handler[model.Album]
}
