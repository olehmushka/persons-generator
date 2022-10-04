package engine

import (
	"context"

	"persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/entities"

	"github.com/google/uuid"
)

type Adapter interface {
	CreateReligions(ctx context.Context, amount int, preferred []*entities.Preference) ([]*religion.Religion, error)
	GetReligionByID(ctx context.Context, id uuid.UUID) (*religion.Religion, error)
	DeleteReligionByID(ctx context.Context, id uuid.UUID) error
	DeleteAllReligions(context.Context) error
}
