package mq

import (
	"context"

	"github.com/google/uuid"
)

type Adapter interface {
	RunAndSaveWorld(context.Context, uuid.UUID, int, int, int, map[uuid.UUID]uuid.UUID) error
	ParseRunAndSaveWorldMsg(context.Context, []byte) (RunAndSaveWorldPayload, error)
}
