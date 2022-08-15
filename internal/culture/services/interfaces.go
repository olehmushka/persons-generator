package services

import (
	"context"

	"persons_generator/internal/culture/entities"
)

type Culture interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error)
}
