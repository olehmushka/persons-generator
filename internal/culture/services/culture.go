package services

import (
	"context"

	"persons_generator/core/storage"
	c "persons_generator/engine/entities/culture"
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

func (s *culture) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*c.SerializedCulture, error) {
	cultures, err := s.engineAdp.CreateCultures(ctx, amount, preferred)
	if err != nil {
		return nil, err
	}

	out := make([]*c.SerializedCulture, len(cultures))
	for i, c := range cultures {
		out[i] = c.Serialize()
	}

	return out, nil
}

func (s *culture) GetProtoCultures(ctx context.Context, q string, opts storage.PaginationSortingOpts) ([]*c.SerializedCulture, int, error) {
	protos, total, err := s.engineAdp.GetProtoCultures(ctx, q, opts)
	if err != nil {
		return nil, 0, err
	}

	out := make([]*c.SerializedCulture, len(protos))
	for i, c := range protos {
		out[i] = c.Serialize()
	}

	return out, total, nil
}

func (s *culture) GetCultureByID(ctx context.Context, id string) (*c.SerializedCulture, error) {
	c, err := s.engineAdp.GetCultureByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return c.Serialize(), nil
}

func (s *culture) DeleteCultureByID(ctx context.Context, id string) error {
	return s.engineAdp.DeleteCultureByID(ctx, id)
}

func (s *culture) DeleteAllCultures(ctx context.Context) error {
	return s.engineAdp.DeleteAllCultures(ctx)
}

func (s *culture) UpdateCultureLanguage(ctx context.Context, cultureID, languageID string) error {
	return s.engineAdp.UpdateCultureLanguage(ctx, cultureID, languageID)
}
