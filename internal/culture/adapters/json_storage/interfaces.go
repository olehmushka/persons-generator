package json_storage

import (
	"context"

	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
)

type Adapter interface {
	SaveCulture(ctx context.Context, c *entities.Culture) error
	GetCultureByID(ctx context.Context, id uuid.UUID) (*entities.Culture, error)
}
