package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	WorldSize      int `env:"WORLD_SIZE"`
	ReligionNumber int `env:"RELIGION_NUMBER"`
}

func New() *Config {
	var cfg Config

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&cfg, "config/default.json"); err != nil {
		fmt.Println(err)
	}

	return &cfg
}
