package services

import (
	"context"

	"persons_generator/internal/religion/entities"
)

type Religion interface {
	CreateReligions(ctx context.Context, amount int, preferences []*entities.Preference) ([]*entities.Religion, error)
	CreateReligionByOriginalCulture(ctx context.Context, oc []byte) (*entities.Religion, error)
}