-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    order_uid          TEXT PRIMARY KEY,
    track_number       TEXT        NOT NULL,
    entry              TEXT        NOT NULL,
    delivery           JSONB       NOT NULL,
    payment            JSONB       NOT NULL,
    items              JSONB       NOT NULL,
    locale             TEXT        NOT NULL,
    internal_signature TEXT        NOT NULL,
    customer_id        TEXT        NOT NULL,
    delivery_service   TEXT        NOT NULL,
    shardkey           TEXT        NOT NULL,
    sm_id              INTEGER     NOT NULL,
    date_created       TIMESTAMPTZ NOT NULL,
    oof_shard          TEXT        NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
