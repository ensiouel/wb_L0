package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/exp/slog"
)

type Config struct {
	Server   Server
	Postgres Postgres
	Nats     Nats
	Logger   Logger
}

type Server struct {
	Addr string `env:"SERVER_ADDR" env-required:"true"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	User     string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DB       string `env:"POSTGRES_DB" env-required:"true"`
}

type Nats struct {
	Addr        string `env:"NATS_ADDR" env-required:"true"`
	ClusterID   string `env:"CLUSTER_ID" env-required:"true"`
	ClientSubID string `env:"CLIENT_SUB_ID" env-required:"true"`
}

type Logger struct {
	Level slog.Level `env:"LOGGER_LEVEL" env-default:"debug"`
}

func New() (Config, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
