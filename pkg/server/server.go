package server

import (
	"context"
	"embed"
	"fmt"
	"future-admin/pkg/env"
	"future-admin/pkg/log"
	"future-admin/pkg/server/handler"
	"future-admin/pkg/utils"
	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	app *fiber.App
}

func NewApiServer() *Server {
	return newServer(Option{
		Name: viper.GetString("name"),
		Cors: &cors.ConfigDefault,
	})
}

func NewWebServer(viewsFs embed.FS) *Server {
	return newServer(Option{
		Name:    viper.GetString("name"),
		ViewDir: "resources/views",
		ViewExt: ".html",
		ViewsFs: viewsFs,
		Cors:    &cors.ConfigDefault,
	})
}

func newServer(opt Option) *Server {
	cfg := fiber.Config{
		AppName:               opt.Name,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler:          handler.ErrorHandler,
		EnablePrintRoutes:     env.IsLocal(),
		DisableStartupMessage: !env.IsLocal(),
	}

	if env.IsLocal() {
		if len(opt.ViewDir) > 0 && len(opt.ViewExt) > 0 {
			cfg.Views = html.New(opt.ViewDir, opt.ViewExt)
		}
	} else {
		cfg.Views = html.NewFileSystem(http.FS(opt.ViewsFs), opt.ViewExt)
	}

	app := fiber.New(cfg)

	app.Use(recover.New(recover.Config{
		EnableStackTrace: !env.IsProduction(),
	}))

	if opt.Cors != nil {
		app.Use(cors.New(*opt.Cors))
	}

	if viper.GetBool("log.stats") {
		app.Use(logger.New(logger.Config{
			DisableColors: true,
			CustomTags: map[string]logger.LogFunc{
				"ip": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
					return output.WriteString(utils.GetLocalIP())
				},
				"latency": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
					latency := data.Stop.Sub(data.Start)
					return output.WriteString(fmt.Sprintf("%v", latency))
				},
			},
			TimeFormat: time.RFC3339,
			Format:     "{\"level\":\"info\",\"ts\":\"${time}\",\"status\":${status},\"duration\":\"${latency}\",\"ip\":\"${ip}\",\"method\":\"${method}\",\"content\":\"${path}\"}\n",
		}))
	}

	if viper.GetBool("pprof.enable") {
		app.Use(pprof.New(pprof.Config{
			Prefix: viper.GetString("pprof.prefix"),
		}))
	}

	app.Use(func(c *fiber.Ctx) error {
		var (
			headers = c.GetReqHeaders()
			values  = make(map[string]string)
		)

		for field, value := range headers {
			if strings.HasPrefix(field, "X-") {
				values[field] = strings.Join(value, "")
			}
		}
		c.SetUserContext(context.WithValue(c.Context(), "x-headers", values))

		return c.Next()
	})

	return &Server{app}
}

func (s *Server) Use(args ...interface{}) {
	s.app.Use(args...)
}

func (s *Server) GET(path string, handlers ...handler.Handler) Router {
	return s.app.Get(path, handler.WrapHandler(handlers...)...)
}

func (s *Server) POST(path string, handlers ...handler.Handler) Router {
	return s.app.Post(path, handler.WrapHandler(handlers...)...)
}

func (s *Server) PUT(path string, handlers ...handler.Handler) Router {
	return s.app.Put(path, handler.WrapHandler(handlers...)...)
}

func (s *Server) PATCH(path string, handlers ...handler.Handler) Router {
	return s.app.Patch(path, handler.WrapHandler(handlers...)...)
}

func (s *Server) DELETE(path string, handlers ...handler.Handler) Router {
	return s.app.Delete(path, handler.WrapHandler(handlers...)...)
}

func (s *Server) View(path string, view string) {
	if !env.IsLocal() {
		view = "resources/views/" + view
	}
	s.app.Get(path, func(c *fiber.Ctx) error {
		return c.Render(view, nil)
	})
}

func (s *Server) Static(fs embed.FS) {
	s.app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(fs),
		PathPrefix: "static",
	}))
}

func (s *Server) StaticDir(dir string) {
	opt := fiber.Static{
		Compress: env.IsProduction(),
		ModifyResponse: func(ctx *fiber.Ctx) error {
			ctx.Response().Header.Set("Access-Control-Allow-Origin", "*")
			return nil
		},
	}
	if !env.IsProduction() {
		opt.CacheDuration = 0
	}
	s.app.Static("/static", dir, opt)
}

func (s *Server) Run() {
	addr := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
	if err := s.app.Listen(addr); err != nil {
		log.Error(err)
	}
}
