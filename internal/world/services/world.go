package services

import (
	"context"
	"persons_generator/internal/world/adapters/engine"
	"persons_generator/internal/world/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type world struct {
	engineAdp engine.Adapter
}

func New(engineAdp engine.Adapter) World {
	return &world{
		engineAdp: engineAdp,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *world) CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*entities.World, error) {
	w, err := s.engineAdp.CreateWorld(ctx, amount, maleAmount, femaleAmount, religionCultureRels)
	if err != nil {
		return nil, err
	}
	if err := w.Save(); err != nil {
		return nil, err
	}

	return serializeWorld(w), nil
}
