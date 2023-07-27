package app

import (
	"context"
	"github.com/ensiouel/wb_L0/internal/config"
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/ensiouel/wb_L0/internal/service"
	"github.com/ensiouel/wb_L0/internal/storage"
	"github.com/ensiouel/wb_L0/internal/transport/http"
	http_handler "github.com/ensiouel/wb_L0/internal/transport/http/handler"
	"github.com/ensiouel/wb_L0/internal/transport/nats"
	nats_handler "github.com/ensiouel/wb_L0/internal/transport/nats/handler"
	"github.com/ensiouel/wb_L0/pkg/postgres"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	conf   config.Config
	logger *slog.Logger
}

func New() *App {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger := initLogger(conf.Logger)

	return &App{
		conf:   conf,
		logger: logger,
	}
}

func (app *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	app.logger.Info("starting app")

	app.logger.Info("connecting to postgres")
	pg, err := initPostgres(ctx, app.conf.Postgres)
	if err != nil {
		app.logger.Error("failed to connect to postgres", slog.Any("error", err))
		return
	}

	cache := ttlcache.New[string, model.Order]()

	orderStorage := storage.NewOrderStorage(pg)
	orderService := service.NewOrderService(orderStorage)

	orderHTTPHandler := http_handler.NewOrderHandler(cache, app.logger)
	orderNATSHandler := nats_handler.NewOrderHandler(app.conf.Nats.ClusterID, orderService, cache, app.logger)

	app.logger.Info("starting preloading cache")
	{
		orders, err := orderService.Select(ctx)
		if err != nil {
			app.logger.Error("failed to select orders", slog.Any("error", err))
			return
		}

		app.logger.Info("preloading orders cache", slog.Any("count", len(orders)))

		for _, order := range orders {
			cache.Set(order.OrderUID, order, ttlcache.NoTTL)
		}
	}

	app.logger.Info("starting http server")
	httpServer := http.New(app.conf.Server, app.logger).Handle(orderHTTPHandler)
	go func() {
		err = httpServer.Run()
		if err != nil {
			app.logger.Error("failed to run http server", slog.Any("error", err))
			return
		}
	}()

	app.logger.Info("starting nats server")
	natsServer, err := nats.New(app.conf.Nats, app.logger)
	if err != nil {
		app.logger.Error("failed to run nats server", slog.Any("error", err))
		return
	}

	err = natsServer.Handle(orderNATSHandler)
	if err != nil {
		app.logger.Error("failed to register nats handlers", slog.Any("error", err))
		return
	}

	<-ctx.Done()

	app.logger.Info("graceful shutdown")

	app.logger.Info("shutting down http server")
	httpServer.Shutdown()

	app.logger.Info("shutting down nats server")
	natsServer.Close()

	app.logger.Info("shutting down postgres")
	pg.Close()
}

func initPostgres(ctx context.Context, conf config.Postgres) (postgres.Client, error) {
	client, err := postgres.NewClient(ctx, postgres.Config{
		Host:     conf.Host,
		Port:     conf.Port,
		User:     conf.User,
		Password: conf.Password,
		DB:       conf.DB,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func initLogger(conf config.Logger) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: conf.Level,
	}))

	return logger
}
