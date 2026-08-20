package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/internal/persons/adapters/engine"
	"persons_generator/internal/persons/entities"

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

func (s *persons) GetPersonsByWorldID(ctx context.Context, worldID string, offset, limit int) ([]*entities.Person, int, error) {
	opts := storage.PaginationSortingOpts{
		Pagination: &storage.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}

	persons, err := s.engineAdp.ReadPersonsByWorldID(ctx, worldID, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not read persons by world id")
	}
	serializedPersons, err := serializeSerializedPeople(persons)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not serialize persons")
	}

	count, err := s.engineAdp.CountPersonsByWorldID(ctx, worldID)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not count persons by world id")
	}

	return serializedPersons, count, nil
}

func (s *persons) DeletePersonByID(ctx context.Context, id string) error {
	return s.engineAdp.DeletePersonByID(ctx, id)
}

func (s *persons) DeleteAllPersons(ctx context.Context) error {
	return s.engineAdp.DeleteAllPersons(ctx)
}
