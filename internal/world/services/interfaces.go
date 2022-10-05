package services

import (
	"context"
	"persons_generator/core/storage"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/adapters/mq"
	"persons_generator/internal/world/entities"

	"github.com/google/uuid"
)

type World interface {
	CreateWorld(context.Context, int, int, int, int, map[uuid.UUID]uuid.UUID) (uuid.UUID, error)
	RunAndSaveWorld(context.Context, []byte) error
	GetWorldRunningProgress(worldID uuid.UUID) (engineWorld.ProgressRunWorld, error)
	ParseRunAndSaveWorldMsg(context.Context, []byte) (mq.RunAndSaveWorldPayload, error)
	DeleteWorldByID(ctx context.Context, id uuid.UUID) error
	DeleteAllWorlds(context.Context) error
	ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.World, int, error)
}
