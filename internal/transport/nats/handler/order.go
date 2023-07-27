package handler

import (
	"context"
	"encoding/json"
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/jellydator/ttlcache/v3"
	"github.com/nats-io/stan.go"
	"golang.org/x/exp/slog"
)

type OrderService interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, uid string) (model.Order, error)
}

type OrderHandler struct {
	natsSubject  string
	orderService OrderService
	cache        *ttlcache.Cache[string, model.Order]
	logger       *slog.Logger
}

func NewOrderHandler(natsSubject string, orderService OrderService, cache *ttlcache.Cache[string, model.Order], logger *slog.Logger) *OrderHandler {
	return &OrderHandler{natsSubject: natsSubject, orderService: orderService, cache: cache, logger: logger}
}

func (handler *OrderHandler) Register(conn stan.Conn) error {
	_, err := conn.Subscribe(handler.natsSubject, handler.create)
	if err != nil {
		return apperror.Internal.WithErrorf("failed to subscribe to cluster: %w", err)
	}

	return nil
}

func (handler *OrderHandler) create(msg *stan.Msg) {
	var order model.Order
	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		handler.logger.Warn("failed to unmarshal order", slog.Any("error", err))
		return
	}

	err = handler.orderService.Create(context.Background(), order)
	if err != nil {
		handler.logger.Warn("failed to create order", slog.Any("error", err))
		return
	}

	handler.cache.Set(order.OrderUID, order, ttlcache.NoTTL)

	handler.logger.Info("order created", slog.Any("order_uid", order.OrderUID))
}
