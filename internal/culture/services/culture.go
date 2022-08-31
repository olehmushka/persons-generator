package services

import (
	"context"

	"persons_generator/internal/culture/adapters/engine"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
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
	cultures, err := s.engineAdp.CreateCultures(ctx, amount, preferred)
	if err != nil {
		return nil, err
	}

	out := make([]*entities.Culture, len(cultures))
	for i, c := range cultures {
		if err := c.Save(); err != nil {
			return nil, err
		}
		out[i] = serializeCulture(c)
	}

	return out, nil
}

func (s *culture) GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*entities.Culture, int, error) {
	protos, total, err := s.engineAdp.GetProtoCultures(ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return serializeCultures(protos), total, nil
}

func (s *culture) GetCultureByID(ctx context.Context, id uuid.UUID) (*entities.Culture, error) {
	c, err := s.engineAdp.GetCultureByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return serializeCulture(c), nil
}

func (s *culture) GetOriginalCultureByID(ctx context.Context, id uuid.UUID) ([]byte, error) {
	// c, err := s.storageAdp.GetCultureByID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }

	// return s.engineAdp.GetOriginalCulture(ctx, c)
	return nil, nil
}

func (s *culture) HybridCulture(ctx context.Context, cultures []*entities.Culture) (*entities.Culture, error) {
	return nil, nil
}
