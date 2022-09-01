package services

import (
	"context"
	"fmt"

	"persons_generator/internal/religion/adapters/engine"
	"persons_generator/internal/religion/entities"

	"go.uber.org/fx"
)

type religion struct {
	engineAdp engine.Adapter
}

func New(engineAdp engine.Adapter) Religion {
	return &religion{
		engineAdp: engineAdp,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *religion) CreateReligions(ctx context.Context, amount int, preferences []*entities.Preference) ([]*entities.Religion, error) {
	if amount < len(preferences) {
		return nil, fmt.Errorf("amount (%d) can not be less than preferences number (%d)", amount, len(preferences))
	}

	rs, err := s.engineAdp.CreateReligions(ctx, amount, preferences)
	if err != nil {
		return nil, err
	}

	out := make([]*entities.Religion, len(rs))
	for i, r := range rs {
		if err := r.Save(); err != nil {
			return nil, err
		}
		out[i] = serializeReligion(r)
	}

	return out, nil
}
