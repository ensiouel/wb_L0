package nats

import (
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/wb_L0/internal/config"
	"github.com/ensiouel/wb_L0/internal/transport/nats/handler"
	"github.com/nats-io/stan.go"
	"golang.org/x/exp/slog"
)

type Server struct {
	conf   config.Nats
	logger *slog.Logger
	conn   stan.Conn
}

func New(conf config.Nats, logger *slog.Logger) (*Server, error) {
	conn, err := stan.Connect(conf.ClusterID, conf.ClientSubID, stan.NatsURL(conf.Addr))
	if err != nil {
		return nil, apperror.Internal.WithErrorf("failed to connect to nats: %w", err)
	}

	return &Server{
		conf:   conf,
		logger: logger,
		conn:   conn,
	}, nil
}

func (server *Server) Handle(orderHandler *handler.OrderHandler) error {
	err := orderHandler.Register(server.conn)
	if err != nil {
		return apperror.Internal.WithErrorf("failed to register order handler: %w", err)
	}

	return nil
}

func (server *Server) Close() error {
	return server.conn.Close()
}
