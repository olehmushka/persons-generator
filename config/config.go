package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	Mode string `env:"MODE" default:""`
}

func New() *Config {
	var cfg Config

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&cfg, "config/default.json"); err != nil {
		fmt.Println(err)
	}

	return &cfg
}
