package main

import (
	"math/rand"
	"time"

	"persons_generator/config"
	"persons_generator/orchestrator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	cfg := config.New()
	orchestrator.New(&orchestrator.Config{
		WorldSize: cfg.WorldSize,
		Religion: orchestrator.ReligionConfig{
			Amount: cfg.ReligionNumber,
		},
	}).ShowReligions()
}
