package services

import (
	"context"
	"persons_generator/internal/persons/entities"
)

type Persons interface {
	GetPersonsByWorldID(ctx context.Context, worldID string, offset, limit int) ([]*entities.Person, int, error)
	DeletePersonByID(ctx context.Context, id string) error
	DeleteAllPersons(context.Context) error
}
