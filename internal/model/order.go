package model

import (
	"time"
)

type Order struct {
	OrderUID          string    `db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string    `db:"track_number" json:"track_number" validate:"required"`
	Entry             string    `db:"entry" json:"entry" validate:"required"`
	Delivery          Delivery  `db:"delivery" json:"delivery" validate:"required"`
	Payment           Payment   `db:"payment" json:"payment" validate:"required"`
	Items             []Item    `db:"items" json:"items" validate:"required"`
	Locale            string    `db:"locale" json:"locale" validate:"required"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature"`
	CustomerID        string    `db:"customer_id" json:"customer_id" validate:"required"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service" validate:"required"`
	Shardkey          string    `db:"shardkey" json:"shardkey" validate:"required"`
	SMID              int       `db:"sm_id" json:"sm_id" validate:"required"`
	DateCreated       time.Time `db:"date_created" json:"date_created" validate:"required"`
	OofShard          string    `db:"oof_shard" json:"oof_shard" validate:"required"`
}
