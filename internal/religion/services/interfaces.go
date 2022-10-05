package services

import (
	"context"

	r "persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/entities"
)

type Religion interface {
	CreateReligions(ctx context.Context, amount int, preferences []*entities.Preference) ([]*r.SerializedReligion, error)
	GetReligionByID(ctx context.Context, id string) (*r.SerializedReligion, error)
	DeleteReligionByID(ctx context.Context, id string) error
	DeleteAllReligions(context.Context) error
}
