package engine

import "context"

type Adapter interface {
	DeleteAllPersons(context.Context) error
}
