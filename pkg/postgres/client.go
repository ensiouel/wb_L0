package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type Client interface {
	Get(ctx context.Context, dest interface{}, query string, args ...any) error
	Select(ctx context.Context, dest interface{}, query string, args ...any) error
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
}

type client struct {
	*pgxpool.Pool
}

func NewClient(ctx context.Context, conf Config) (Client, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host,
		conf.Port, conf.DB))
	if err != nil {
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &client{pool}, nil
}

func (client *client) Get(ctx context.Context, dest interface{}, query string, args ...any) error {
	return pgxscan.Get(ctx, client.Pool, dest, query, args...)
}

func (client *client) Select(ctx context.Context, dest interface{}, query string, args ...any) error {
	return pgxscan.Select(ctx, client.Pool, dest, query, args...)
}
