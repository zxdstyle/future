package handler

import (
	"future-admin/internal/logic/album"
	"future-admin/internal/model"
	"future-admin/pkg/base"
	"github.com/samber/do"
)

var Album = &albumHandler{
	Handler: base.NewHandler[model.Album](
		do.MustInvoke[*album.Logic](nil).Logic,
	),
}

type albumHandler struct {
	*base.Handler[model.Album]
}
