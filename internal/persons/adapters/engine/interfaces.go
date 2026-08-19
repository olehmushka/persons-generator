package engine

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/engine/entities/person"
)

type Adapter interface {
	DeletePersonByID(ctx context.Context, id string) error
	DeleteAllPersons(context.Context) error
	ReadPersonsByWorldID(ctx context.Context, worldID string, opts storage.PaginationSortingOpts) ([]*person.SerializedPerson, error)
	CountPersonsByWorldID(ctx context.Context, worldID string) (int, error)
}
