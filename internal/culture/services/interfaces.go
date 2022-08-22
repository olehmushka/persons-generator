package services

import (
	"context"

	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
)

type Culture interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error)
	GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*entities.Culture, int, error)
	GetCultureByID(ctx context.Context, id uuid.UUID) (*entities.Culture, error)
	GetOriginalCultureByID(ctx context.Context, id uuid.UUID) ([]byte, error)
	HybridCulture(ctx context.Context, cultures []*entities.Culture) (*entities.Culture, error)
}
