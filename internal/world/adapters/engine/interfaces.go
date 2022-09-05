package engine

import (
	"context"
	"persons_generator/engine/entities/world"

	"github.com/google/uuid"
)

type Adapter interface {
	CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*world.World, error)
}
