package orchestrator

import (
	"persons_generator/entities"
	"persons_generator/religion"
	"persons_generator/world"

	"github.com/davecgh/go-spew/spew"
)

type Orchestrator struct {
	w         *entities.World
	cultures  []*entities.Culture
	religions []*entities.Religion
}

func New(cfg *Config) *Orchestrator {
	w := world.FillWorld(&entities.World{
		Size: cfg.WorldSize,
	})
	r := religion.NewMany(cfg.ReligionNumber)

	return &Orchestrator{
		w:         w,
		religions: r,
	}
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowReligions() {
	spew.Config.MaxDepth = 10
	for _, r := range o.religions {
		spew.Dump(r)
	}
}
