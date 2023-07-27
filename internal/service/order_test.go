package service_test

import (
	"context"
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/ensiouel/wb_L0/internal/service"
	"github.com/ensiouel/wb_L0/internal/storage/mock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var staticOrder = model.Order{
	OrderUID:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: model.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		ZIP:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: model.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDT:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	},
	Items: []model.Item{
		{
			ChrtID:      9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			RID:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmID:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202,
		},
	},
	Locale:            "en",
	InternalSignature: "",
	CustomerID:        "test",
	DeliveryService:   "meest",
	Shardkey:          "9",
	SMID:              99,
	DateCreated:       time.Date(2021, 11, 1, 26, 22, 19, 0, time.UTC),
	OofShard:          "1",
}

func TestOrderService_Create(t *testing.T) {
	var cases = []struct {
		name    string
		mock    *storage.OrderStorageMock
		in      model.Order
		err     error
		wantErr bool
	}{
		{
			name: "ok",
			mock: &storage.OrderStorageMock{
				CreateFunc: func(ctx context.Context, order model.Order) error {
					return nil
				},
			},
			in:      staticOrder,
			err:     nil,
			wantErr: false,
		},
		{
			name: "bad request",
			mock: &storage.OrderStorageMock{
				CreateFunc: func(ctx context.Context, order model.Order) error {
					return nil
				},
			},
			in:      model.Order{},
			err:     apperror.BadRequest,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			orderService := service.NewOrderService(tc.mock)

			err := orderService.Create(context.Background(), tc.in)
			if tc.wantErr {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOrderService_Get(t *testing.T) {
	var cases = []struct {
		name    string
		mock    *storage.OrderStorageMock
		in      string
		out     model.Order
		err     error
		wantErr bool
	}{
		{
			name: "ok",
			mock: &storage.OrderStorageMock{
				GetFunc: func(ctx context.Context, orderUID string) (model.Order, error) {
					return staticOrder, nil
				},
			},
			in:      staticOrder.OrderUID,
			out:     staticOrder,
			err:     nil,
			wantErr: false,
		},
		{
			name: "not found",
			mock: &storage.OrderStorageMock{
				GetFunc: func(ctx context.Context, orderUID string) (model.Order, error) {
					return model.Order{}, apperror.NotFound
				},
			},
			in:      "abcd",
			err:     apperror.NotFound,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			orderService := service.NewOrderService(tc.mock)

			order, err := orderService.Get(context.Background(), tc.in)
			if tc.wantErr {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.Equal(t, tc.out, order)
			}
		})
	}
}

func TestOrderService_Select(t *testing.T) {
	staticOrders := []model.Order{staticOrder}

	var cases = []struct {
		name    string
		mock    *storage.OrderStorageMock
		out     []model.Order
		err     error
		wantErr bool
	}{
		{
			name: "ok",
			mock: &storage.OrderStorageMock{
				SelectFunc: func(ctx context.Context) ([]model.Order, error) {
					return staticOrders, nil
				},
			},
			out:     staticOrders,
			err:     nil,
			wantErr: false,
		},
		{
			name: "not found",
			mock: &storage.OrderStorageMock{
				SelectFunc: func(ctx context.Context) ([]model.Order, error) {
					return []model.Order{}, apperror.NotFound
				},
			},
			err:     apperror.NotFound,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			orderService := service.NewOrderService(tc.mock)

			orders, err := orderService.Select(context.Background())
			if tc.wantErr {
				assert.ErrorIs(t, err, tc.err)
			} else {
				assert.Equal(t, tc.out, orders)
			}
		})
	}
}
