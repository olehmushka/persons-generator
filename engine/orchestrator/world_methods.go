package orchestrator

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
)

func (o *Orchestrator) CreateWorld(
	size int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*world.World, error) {
	return world.New(world.Config{StorageFolderName: o.storageFolderName}, size, cultures, religions, refs)
}
