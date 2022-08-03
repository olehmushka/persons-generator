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
	c := culture.NewCultures(cfg.Culture.Amount, cfg.Culture.Preferred)
	r := religion.NewReligions(culture.GetReligionCultures(cfg.Religion.Amount, c))

	return &Orchestrator{
		w:         w,
		cultures:  c,
		religions: r,
	}
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowCultures() {
	for _, c := range o.cultures {
		c.Print()
	}
}

func (o *Orchestrator) ShowReligions() {
	for _, r := range o.religions {
		r.Print()
	}
}
