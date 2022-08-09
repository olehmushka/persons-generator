package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	l "persons_generator/entities/language"
	"persons_generator/orchestrator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if err := l.SetWordBases(); err != nil {
		fmt.Printf("[language.SetWordBases] error = %+v\n", err)
		os.Exit(1)
		return
	}
	// cfg := config.New()
	o, err := orchestrator.New(&orchestrator.Config{
		WorldSize: 5,
		Culture: orchestrator.CultureConfig{
			Preferred: []string{"indo_aryan_bengali"},
			Amount:    1,
		},
		Religion: orchestrator.ReligionConfig{
			Amount: 1,
		},
	})
	if err != nil {
		fmt.Printf("[Orchestrator.New] error = %+v\n", err)
		os.Exit(1)
		return
	}
	o.ShowCultures()
	// o.ShowReligions()
}
