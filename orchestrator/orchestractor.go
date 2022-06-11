package orchestrator

import (
	"persons_generator/entities"
	newReligion "persons_generator/entities/religion/religion"
	"persons_generator/world"

	"github.com/davecgh/go-spew/spew"
)

type Orchestrator struct {
	w         *entities.World
	cultures  []*entities.Culture
	religions []*newReligion.Religion
}

func New(cfg *Config) *Orchestrator {
	w := world.FillWorld(&entities.World{
		Size: cfg.WorldSize,
	})
	r := newReligion.NewReligions(cfg.ReligionNumber)

	return &Orchestrator{
		w:         w,
		religions: r,
	}
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowReligions() {
	spew.Config.MaxDepth = 10
	spew.Config.DisableMethods = true
	for _, r := range o.religions {
		spew.Dump(r)
	}
}
