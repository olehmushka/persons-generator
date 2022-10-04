package services

import (
	"context"
	"persons_generator/internal/persons/adapters/engine"
	"persons_generator/internal/persons/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type persons struct {
	engineAdp engine.Adapter
}

func New(engineAdp engine.Adapter) Persons {
	return &persons{engineAdp: engineAdp}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *persons) GetPersonsByWorldID(ctx context.Context, worldID uuid.UUID, offset, limit int) ([]*entities.Person, int, error) {
	return nil, 0, nil
}

func (s *persons) DeleteAllPersons(ctx context.Context) error {
	return s.engineAdp.DeleteAllPersons(ctx)
}
