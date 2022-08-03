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
	o := orchestrator.New(&orchestrator.Config{
		WorldSize: cfg.WorldSize,
		Culture: orchestrator.CultureConfig{
			Preferred: []string{"ruthenian"},
			Amount:    1,
		},
		Religion: orchestrator.ReligionConfig{
			Amount: cfg.ReligionNumber,
		},
	})
	// o.ShowReligions()
	o.ShowCultures()
}
