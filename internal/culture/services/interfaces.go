package services

import (
	"context"

	c "persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
)

type Culture interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*c.SerializedCulture, error)
	GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*c.SerializedCulture, int, error)
	GetCultureByID(ctx context.Context, id uuid.UUID) (*c.SerializedCulture, error)
	DeleteCultureByID(ctx context.Context, id uuid.UUID) error
	DeleteAllCultures(context.Context) error
}
