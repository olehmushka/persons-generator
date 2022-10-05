package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/internal/language/adapters/engine"
	"persons_generator/internal/language/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type persons struct {
	engineAdp engine.Adapter
}

func New(engineAdp engine.Adapter) Language {
	return &persons{engineAdp: engineAdp}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *persons) QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error) {
	langs, err := s.engineAdp.QueryDefaultLanguages(q, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not query default languages")
	}
	count, err := s.engineAdp.CountDefaultLanguages(q)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not count default languages")
	}

	return langs, count, nil
}

func (s *persons) CreateLanguage(ctx context.Context, in *entities.Language) (uuid.UUID, error) {
	return s.engineAdp.CreateLanguage(ctx, in)
}

func (s *persons) ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error) {
	langs, err := s.engineAdp.ReadLanguagesByName(ctx, name, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not read languages by name")
	}
	count, err := s.engineAdp.CountLanguagesByName(ctx, name)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not count languages by name")
	}

	return langs, count, nil
}

func (s *persons) DeleteLanguageByID(ctx context.Context, id uuid.UUID) error {
	return s.engineAdp.DeleteLanguageByID(ctx, id)
}

func (s *persons) DeleteAllLanguages(ctx context.Context) error {
	return s.engineAdp.DeleteAllLanguages(ctx)
}
