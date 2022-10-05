package services

import (
	"context"

	c "persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"
)

type Culture interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*c.SerializedCulture, error)
	GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*c.SerializedCulture, int, error)
	GetCultureByID(ctx context.Context, id string) (*c.SerializedCulture, error)
	DeleteCultureByID(ctx context.Context, id string) error
	DeleteAllCultures(context.Context) error
}
