package engine

import (
	"context"
)

type Adapter interface {
	DeletePersonByID(ctx context.Context, id string) error
	DeleteAllPersons(context.Context) error
}
