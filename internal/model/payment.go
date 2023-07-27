package model

type Payment struct {
	Transaction  string `db:"transaction" json:"transaction" validate:"required"`
	RequestID    string `db:"request_id" json:"request_id"`
	Currency     string `db:"currency" json:"currency" validate:"required"`
	Provider     string `db:"provider" json:"provider" validate:"required"`
	Amount       int    `db:"amount" json:"amount" validate:"gt=0"`
	PaymentDT    int64  `db:"payment_dt" json:"payment_dt" validate:"required"`
	Bank         string `db:"bank" json:"bank" validate:"required"`
	DeliveryCost int    `db:"delivery_cost" json:"delivery_cost" validate:"gt=0"`
	GoodsTotal   int    `db:"goods_total" json:"goods_total" validate:"gt=0"`
	CustomFee    int    `db:"custom_fee" json:"custom_fee" validate:"gte=0"`
}
