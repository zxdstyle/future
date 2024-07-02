package server

import (
	"embed"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Option struct {
	Name    string
	ViewDir string
	ViewExt string
	ViewsFs embed.FS
	Cors    *cors.Config
}
