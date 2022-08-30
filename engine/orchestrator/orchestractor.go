package orchestrator

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
)

type Orchestrator struct {
	w         *world.World
	cultures  []*culture.Culture
	religions []*religion.Religion

	storageFolderName string
}

func New(cfg Config) (*Orchestrator, error) {
	return &Orchestrator{
		storageFolderName: cfg.StorageFolderName,
	}, nil
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
