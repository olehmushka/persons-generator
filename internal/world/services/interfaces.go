package services

import (
	"context"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/adapters/mq"

	"github.com/google/uuid"
)

type World interface {
	CreateWorld(context.Context, int, int, int, map[uuid.UUID]uuid.UUID) (uuid.UUID, error)
	RunAndSaveWorld(context.Context, []byte) error
	GetWorldRunningProgress(worldID uuid.UUID) (engineWorld.ProgressRunWorld, error)
	ParseRunAndSaveWorldMsg(context.Context, []byte) (mq.RunAndSaveWorldPayload, error)
}
