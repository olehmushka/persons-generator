package services

import (
	"context"

	"persons_generator/internal/culture/adapters/engine"
	"persons_generator/internal/culture/entities"

	"go.uber.org/fx"
)

type culture struct {
	engineAdp engine.Adapter
}

func New(engineAdp engine.Adapter) Culture {
	return &culture{
		engineAdp: engineAdp,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *culture) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error) {
	return s.engineAdp.CreateCultures(ctx, amount, preferred)
}
