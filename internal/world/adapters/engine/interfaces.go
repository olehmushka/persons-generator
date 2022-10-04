package engine

import (
	"context"
	"persons_generator/engine/entities/world"

	"github.com/google/uuid"
)

type Adapter interface {
	CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*world.World, error)
	RunAndSaveWorld(ctx context.Context, w *world.World, stopYear int) error
	GetWorldRunningProgress(worldID uuid.UUID) (world.ProgressRunWorld, error)
	DeleteAllWorlds(context.Context) error
}
