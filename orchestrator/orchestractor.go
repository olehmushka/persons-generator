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
	cultures := culture.NewCultures(cfg.Culture.Amount, cfg.Culture.Preferred)
	w = w.CulturesPropagate(cultures)
	locationsWithReligion := w.GetLocationsForReligionGenerate(cfg.Religion.Amount)
	locationsWithReligion = world.FillLocationsWithReligions(locationsWithReligion)
	religions := world.ExtractReligionFromLocations(locationsWithReligion)
	w = w.ReligionsPropagate(locationsWithReligion)

	return &Orchestrator{
		w:         w,
		cultures:  cultures,
		religions: religions,
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
