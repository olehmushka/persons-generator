package engine

import (
	"context"

	"persons_generator/internal/culture/entities"
)

type Adapter interface {
	CreateCultures(ctx context.Context, amount int, preferred []string) ([]*entities.Culture, error)
}
