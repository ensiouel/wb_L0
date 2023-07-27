package storage

import (
	"context"
	"errors"
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/wb_L0/internal/model"
	"github.com/ensiouel/wb_L0/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type OrderStorage interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, uid string) (model.Order, error)
	Select(ctx context.Context) ([]model.Order, error)
}

type OrderStorageImpl struct {
	client postgres.Client
}

func NewOrderStorage(client postgres.Client) *OrderStorageImpl {
	return &OrderStorageImpl{client: client}
}

func (storage *OrderStorageImpl) Create(ctx context.Context, order model.Order) error {
	q := `
INSERT INTO orders (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id,
                    delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
`

	_, err := storage.client.Exec(ctx, q,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Delivery,
		order.Payment,
		order.Items,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SMID,
		order.DateCreated,
		order.OofShard,
	)
	if err != nil {
		return apperror.Internal.WithError(err)
	}

	return nil
}

func (storage *OrderStorageImpl) Get(ctx context.Context, uid string) (model.Order, error) {
	q := `
SELECT order_uid,
       track_number,
       entry,
       delivery,
       payment,
       items,
       locale,
       internal_signature,
       customer_id,
       delivery_service,
       shardkey,
       sm_id,
       date_created,
       oof_shard
FROM orders
WHERE order_uid = $1;
`

	var order model.Order
	err := storage.client.Get(ctx, &order, q, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, apperror.NotFound.WithError(err)
		}

		return model.Order{}, apperror.Internal.WithError(err)
	}

	return order, nil
}

func (storage *OrderStorageImpl) Select(ctx context.Context) ([]model.Order, error) {
	q := `
SELECT order_uid,
       track_number,
       entry,
       delivery,
       payment,
       items,
       locale,
       internal_signature,
       customer_id,
       delivery_service,
       shardkey,
       sm_id,
       date_created,
       oof_shard
FROM orders;
`

	var orders []model.Order
	err := storage.client.Select(ctx, &orders, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.Order{}, apperror.NotFound.WithError(err)
		}

		return []model.Order{}, apperror.Internal.WithError(err)
	}

	return orders, nil
}
