package engine

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/engine/entities/world"

	"github.com/google/uuid"
)

type Adapter interface {
	CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*world.World, error)
	RunAndSaveWorld(ctx context.Context, w *world.World, stopYear int) error
	GetWorldRunningProgress(worldID uuid.UUID) (world.ProgressRunWorld, error)
	DeleteWorldByID(ctx context.Context, id uuid.UUID) error
	DeleteAllWorlds(context.Context) error
	ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*world.SerializedWorld, error)
	CountWorlds(ctx context.Context, opts storage.PaginationSortingOpts) (int, error)
}
