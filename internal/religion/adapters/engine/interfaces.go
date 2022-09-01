package engine

import (
	"context"

	"persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/entities"
)

type Adapter interface {
	CreateReligions(ctx context.Context, amount int, preferred []*entities.Preference) ([]*religion.Religion, error)
}
