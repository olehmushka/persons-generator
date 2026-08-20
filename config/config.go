package config

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Config struct {
	HTTPServer HTTPServer
	MongoDB    MongoDB
	Redis      Redis
}

func New() (*Config, error) {
	var cfg Config

	// A missing .env is expected outside local development (containers and
	// CI inject env vars directly) and shouldn't fail startup; a malformed
	// .env that exists but can't be parsed should still be surfaced.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&cfg, "config/default.json"); err != nil {
		return nil, err
	}

	return &cfg, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
