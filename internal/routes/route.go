package routes

import (
	"future-admin/internal/handler"
	"future-admin/pkg/server"
)

func RegisterRoute(srv *server.Server) {

	srv.GET("/api/v1/drivers", handler.Driver.List)

	srv.GET("/api/v1/fs", handler.Fs.List)
	srv.GET("/api/v1/fs/:path", handler.Fs.Detail)
	srv.GET("/api/v1/fs/:file/preview", handler.Fs.Preview)

	srv.GET("/api/v1/albums", handler.Album.List)
	srv.POST("/api/v1/albums", handler.Album.Create)
	srv.GET("/api/v1/albums/:id", handler.Album.Show)
	srv.PUT("/api/v1/albums/:id", handler.Album.Update)
	srv.POST("/api/v1/albums", handler.Album.Create)
	srv.DELETE("/api/v1/albums/:id", handler.Album.Destroy)

	srv.GET("/api/v1/storages", handler.Storage.List)
	srv.POST("/api/v1/storages", handler.Storage.Create)
	srv.GET("/api/v1/storages/:id", handler.Storage.Show)
	srv.PUT("/api/v1/storages/:id", handler.Storage.Update)
	srv.POST("/api/v1/storages", handler.Storage.Create)
	srv.DELETE("/api/v1/storages/:id", handler.Storage.Destroy)

	srv.POST("/api/v1/images", handler.Image.Create)
}
