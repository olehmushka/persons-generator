package engine

import (
	"context"

	"persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
)

type Adapter interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*culture.Culture, error)
	GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*culture.Culture, int, error)
	GetCultureByID(ctx context.Context, id uuid.UUID) (*culture.Culture, error)
	DeleteAllCultures(context.Context) error
}
