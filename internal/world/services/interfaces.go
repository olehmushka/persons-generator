package services

import (
	"context"
	"persons_generator/internal/world/entities"

	"github.com/google/uuid"
)

type World interface {
	CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*entities.World, error)
}
