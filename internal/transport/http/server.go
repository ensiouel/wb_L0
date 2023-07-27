package http

import (
	"github.com/ensiouel/wb_L0/internal/config"
	"github.com/ensiouel/wb_L0/internal/transport/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/exp/slog"
)

type Server struct {
	conf   config.Server
	logger *slog.Logger
	app    *fiber.App
}

func New(conf config.Server, logger *slog.Logger) *Server {
	views := html.New("static/template", ".tmpl")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
	})

	return &Server{
		conf:   conf,
		logger: logger,
		app:    app,
	}
}

func (server *Server) Handle(orderHandler *handler.OrderHandler) *Server {
	orderHandler.Register(server.app.Group("order"))
	return server
}

func (server *Server) Run() error {
	return server.app.Listen(server.conf.Addr)
}

func (server *Server) Shutdown() error {
	return server.app.Shutdown()
}
