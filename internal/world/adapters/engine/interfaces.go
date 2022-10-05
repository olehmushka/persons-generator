package engine

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/engine/entities/world"
)

type Adapter interface {
	CreateWorld(ctx context.Context, id string, amount, maleAmount, femaleAmount int, religionCultureRels map[string]string) (*world.World, error)
	RunAndSaveWorld(ctx context.Context, w *world.World, stopYear int) error
	GetWorldRunningProgress(worldID string) (world.ProgressRunWorld, error)
	DeleteWorldByID(ctx context.Context, id string) error
	DeleteAllWorlds(context.Context) error
	ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*world.SerializedWorld, error)
	CountWorlds(ctx context.Context, opts storage.PaginationSortingOpts) (int, error)
}
