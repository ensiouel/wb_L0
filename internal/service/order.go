package service

import (
	"context"
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/ensiouel/wb_L0/internal/storage"
	"github.com/go-playground/validator/v10"
)

type OrderService interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, uid string) (model.Order, error)
}

type OrderServiceImpl struct {
	orderStorage storage.OrderStorage
	validate     *validator.Validate
}

func NewOrderService(orderStorage storage.OrderStorage) *OrderServiceImpl {
	validate := validator.New()

	return &OrderServiceImpl{orderStorage: orderStorage, validate: validate}
}

func (service *OrderServiceImpl) Create(ctx context.Context, order model.Order) error {
	if err := service.validate.Struct(&order); err != nil {
		return apperror.BadRequest.WithError(err)
	}

	err := service.orderStorage.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (service *OrderServiceImpl) Get(ctx context.Context, uid string) (model.Order, error) {
	order, err := service.orderStorage.Get(ctx, uid)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (service *OrderServiceImpl) Select(ctx context.Context) ([]model.Order, error) {
	orders, err := service.orderStorage.Select(ctx)
	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
