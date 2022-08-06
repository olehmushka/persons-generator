package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"persons_generator/orchestrator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// cfg := config.New()
	o, err := orchestrator.New(&orchestrator.Config{
		WorldSize: 8,
		Culture: orchestrator.CultureConfig{
			Preferred: []string{"ruthenian", "polish"},
			Amount:    5,
		},
		Religion: orchestrator.ReligionConfig{
			Amount: 2,
		},
	})
	if err != nil {
		fmt.Printf("[Orchestrator.New] error = %+v\n", err)
		os.Exit(1)
		return
	}
	o.ShowCultures()
	o.ShowReligions()
}
