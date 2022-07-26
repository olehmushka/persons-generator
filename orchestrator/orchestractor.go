package orchestrator

import (
	"persons_generator/entities/culture"
	"persons_generator/entities/religion"
	"persons_generator/entities/world"
)

type Orchestrator struct {
	w         *world.World
	cultures  []*culture.Culture
	religions []*religion.Religion
}

func New(cfg *Config) *Orchestrator {
	w := world.New(cfg.WorldSize).Fill()
	r := religion.NewReligions(cfg.Religion.Amount)

	return &Orchestrator{
		w:         w,
		religions: r,
	}
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowReligions() {
	for _, r := range o.religions {
		r.Print()
	}
}
