package services

import (
	"context"

	"persons_generator/internal/culture/adapters/engine"
	js "persons_generator/internal/culture/adapters/json_storage"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type culture struct {
	engineAdp  engine.Adapter
	storageAdp js.Adapter
}

func New(engineAdp engine.Adapter, storageAdp js.Adapter) Culture {
	return &culture{
		engineAdp:  engineAdp,
		storageAdp: storageAdp,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *culture) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error) {
	cultures, err := s.engineAdp.CreateCultures(ctx, amount, preferred)
	if err != nil {
		return nil, err
	}
	out := make([]*entities.Culture, len(cultures))
	for i, c := range cultures {
		c.ID = uuid.New()
		if err := s.storageAdp.SaveCulture(ctx, c); err != nil {
			return nil, err
		}
		out[i] = c
	}

	return out, nil
}

func (s *culture) GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*entities.Culture, int, error) {
	return s.engineAdp.GetProtoCultures(ctx, q, limit, offset)
}

func (s *culture) GetCultureByID(ctx context.Context, id uuid.UUID) (*entities.Culture, error) {
	return s.storageAdp.GetCultureByID(ctx, id)
}
