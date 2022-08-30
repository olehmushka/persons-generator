package engine

import (
	"context"

	"persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"
)

type Adapter interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*culture.Culture, error)
	GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*culture.Culture, int, error)
	GetOriginalCulture(ctx context.Context, c *entities.Culture) ([]byte, error)
	HybridCulture(ctx context.Context, cultures []*entities.Culture) (*culture.Culture, error)
}
