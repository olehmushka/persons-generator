package services

import (
	"context"
	"persons_generator/core/storage"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/adapters/mq"
	"persons_generator/internal/world/entities"
)

type World interface {
	CreateWorld(context.Context, int, int, int, int, map[string]string) (string, error)
	RunAndSaveWorld(context.Context, []byte) error
	GetWorldRunningProgress(worldID string) (engineWorld.ProgressRunWorld, error)
	ParseRunAndSaveWorldMsg(context.Context, []byte) (*mq.RunAndSaveWorldPayload, error)
	DeleteWorldByID(ctx context.Context, id string) error
	DeleteAllWorlds(context.Context) error
	ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.World, int, error)
}
