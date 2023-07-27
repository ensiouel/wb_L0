package handler

import (
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/exp/slog"
)

type OrderHandler struct {
	cache  *ttlcache.Cache[string, model.Order]
	logger *slog.Logger
}

func NewOrderHandler(cache *ttlcache.Cache[string, model.Order], logger *slog.Logger) *OrderHandler {
	return &OrderHandler{cache: cache, logger: logger}
}

func (handler *OrderHandler) Register(router fiber.Router) {
	router.Get("", handler.get)
}

func (handler *OrderHandler) get(c *fiber.Ctx) error {
	uid := c.Query("uid")
	if uid == "" {
		return c.Render("index", fiber.Map{})
	}

	order := handler.cache.Get(uid)
	if order == nil {
		return c.Render("index", fiber.Map{})
	}

	return c.Render("index", fiber.Map{
		"order": order.Value(),
	})
}
