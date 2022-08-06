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

func New(cfg *Config) (*Orchestrator, error) {
	w := world.New(cfg.WorldSize).Fill()
	var err error
	w, err = w.CulturesPropagate(cfg.Culture.Amount, cfg.Culture.Preferred)
	if err != nil {
		return nil, err
	}
	w, err = w.ReligionsPropagate(cfg.Religion.Amount)
	if err != nil {
		return nil, err
	}

	return &Orchestrator{
		w:         w,
		cultures:  w.Cultures,
		religions: w.Religions,
	}, nil
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowCultures() {
	for _, c := range o.cultures {
		c.Print()
	}
	o.w.PrintLocationCultures()
}

func (o *Orchestrator) ShowReligions() {
	for _, r := range o.religions {
		r.Print()
	}
	o.w.PrintLocationReligions()
}
