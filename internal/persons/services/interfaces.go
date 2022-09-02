package services

import (
	"context"
	"persons_generator/internal/persons/entities"

	"github.com/google/uuid"
)

type Persons interface {
	GetPersonsByWorldID(ctx context.Context, worldID uuid.UUID, offset, limit int) ([]*entities.Person, int, error)
}
