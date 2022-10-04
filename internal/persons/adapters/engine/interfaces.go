package engine

import (
	"context"

	"github.com/google/uuid"
)

type Adapter interface {
	DeletePersonByID(ctx context.Context, id uuid.UUID) error
	DeleteAllPersons(context.Context) error
}
