package routes

import (
	"future-admin/internal/handler"
	"future-admin/pkg/server"
)

func RegisterRoute(srv *server.Server) {
	srv.GET("/api/v1/storage", handler.Storage.List)
	srv.GET("/api/v1/storage/:path", handler.Storage.Detail)
	srv.GET("/api/v1/storage/:file/preview", handler.Storage.Preview)

	srv.GET("/api/v1/albums", handler.Album.List)
	srv.POST("/api/v1/albums", handler.Album.Create)
	srv.POST("/api/v1/albums/:album", handler.Album.Update)
	srv.POST("/api/v1/albums", handler.Album.Create)
}
