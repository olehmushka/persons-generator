package services

import (
	"context"
	"persons_generator/core/storage"
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

func (s *persons) QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	return s.engineAdp.QueryDefaultLanguages(q, opts)
}

func (s *persons) CreateLanguage(ctx context.Context, in *entities.Language) (uuid.UUID, error) {
	return s.engineAdp.CreateLanguage(ctx, in)
}

func (s *persons) ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	return s.engineAdp.ReadLanguagesByName(ctx, name, opts)
}

func (s *persons) DeleteLanguageByID(ctx context.Context, id uuid.UUID) error {
	return s.engineAdp.DeleteLanguageByID(ctx, id)
}

func (s *persons) DeleteAllLanguages(ctx context.Context) error {
	return s.engineAdp.DeleteAllLanguages(ctx)
}
