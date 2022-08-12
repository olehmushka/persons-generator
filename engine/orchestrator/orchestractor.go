package orchestrator

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
)

type Orchestrator struct {
	w         *world.World
	cultures  []*culture.Culture
	religions []*religion.Religion
}

func New() (*Orchestrator, error) {
	return &Orchestrator{}, nil
	// w := world.New(cfg.WorldSize).Fill()
	// var err error
	// w, err = w.CulturesPropagate(cfg.Culture.Amount, cfg.Culture.Preferred)
	// if err != nil {
	// 	return nil, err
	// }
	// w, err = w.ReligionsPropagate(cfg.Religion.Amount)
	// if err != nil {
	// 	return nil, err
	// }

	// return &Orchestrator{
	// 	w:         w,
	// 	cultures:  w.Cultures,
	// 	religions: w.Religions,
	// }, nil
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowCultures() {
	fmt.Println()
	for _, c := range o.cultures {
		c.Print()
	}
	o.w.PrintLocationCultures()
	fmt.Println()
}

func (o *Orchestrator) ShowReligions() {
	fmt.Println()
	for _, r := range o.religions {
		r.Print()
	}
	o.w.PrintLocationReligions()
	fmt.Println()
}

func (o *Orchestrator) CreateCultures(amount int, preferred []string) ([]*culture.Culture, error) {
	return culture.NewCultures(amount, preferred)
}
