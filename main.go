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
		WorldSize: 5,
		Culture: orchestrator.CultureConfig{
			Preferred: []string{},
			Amount:    4,
		},
		Religion: orchestrator.ReligionConfig{
			Amount: 10,
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
