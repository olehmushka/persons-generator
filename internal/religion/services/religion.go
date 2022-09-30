package services

import (
	"context"
	"fmt"

	"persons_generator/core/wrapped_error"
	r "persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/adapters/engine"
	"persons_generator/internal/religion/entities"

	"github.com/google/uuid"
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

func (s *religion) CreateReligions(ctx context.Context, amount int, preferences []*entities.Preference) ([]*r.SerializedReligion, error) {
	if amount < len(preferences) {
		return nil, wrapped_error.NewBadRequestError(nil, fmt.Sprintf("amount (%d) can not be less than preferences number (%d)", amount, len(preferences)))
	}

	rs, err := s.engineAdp.CreateReligions(ctx, amount, preferences)
	if err != nil {
		return nil, err
	}

	out := make([]*r.SerializedReligion, len(rs))
	for i, r := range rs {
		if err := r.Save(); err != nil {
			return nil, err
		}
		out[i] = r.Serialize()
	}

	return out, nil
}

func (s *religion) GetReligionByID(ctx context.Context, id uuid.UUID) (*r.SerializedReligion, error) {
	relig, err := s.engineAdp.GetReligionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return relig.Serialize(), nil
}
