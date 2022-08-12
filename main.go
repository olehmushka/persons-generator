package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"persons_generator/entities/human"
	l "persons_generator/entities/language"
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
	// o, err := orchestrator.New(&orchestrator.Config{
	// 	WorldSize: 5,
	// 	Culture: orchestrator.CultureConfig{
	// 		Preferred: []string{"lith"},
	// 		Amount:    1,
	// 	},
	// 	Religion: orchestrator.ReligionConfig{
	// 		Amount: 1,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("[Orchestrator.New] error = %+v\n", err)
	// 	os.Exit(1)
	// 	return
	// }
	// o.ShowCultures()
	// o.ShowReligions()

	v, err := human.GenerateHuman()
	if err != nil {
		fmt.Printf("[Orchestrator.New.GenerateHuman] error = %+v\n", err)
	}
	v.Print()
}
