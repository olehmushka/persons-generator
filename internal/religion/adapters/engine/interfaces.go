package engine

import (
	"context"

	cultureEntities "persons_generator/internal/culture/entities"
	"persons_generator/internal/religion/entities"
)

type Adapter interface {
	CreateReligions(ctx context.Context, cultures []*cultureEntities.Culture) ([]*entities.Religion, error)
}
