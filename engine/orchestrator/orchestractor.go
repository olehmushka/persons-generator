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
}

func (o *Orchestrator) Orchestrate() {
}
