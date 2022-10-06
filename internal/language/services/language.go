package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/internal/language/adapters/engine"
	"persons_generator/internal/language/entities"

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

func (s *persons) CreateLanguage(ctx context.Context, in *entities.Language) (string, error) {
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

func (s *persons) ReadDefaultLanguageSubfamilies(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.Subfamily, int, error) {
	sfs, err := s.engineAdp.ReadDefaultLanguageSubfamilies(ctx, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not read default language subfamilies")
	}
	count, err := s.engineAdp.CountDefaultLanguageSubfamilies(ctx)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not count default language subfamilies")
	}

	return sfs, count, nil
}

func (s *persons) ReadLanguage(ctx context.Context, id, name string) (*entities.Language, error) {
	return nil, nil
}

func (s *persons) DeleteLanguageByID(ctx context.Context, id string) error {
	return s.engineAdp.DeleteLanguageByID(ctx, id)
}

func (s *persons) DeleteAllLanguages(ctx context.Context) error {
	return s.engineAdp.DeleteAllLanguages(ctx)
}

func (s *persons) ReadLanguageByID(ctx context.Context, id string) (*entities.Language, error) {
	return s.engineAdp.ReadLanguageByID(ctx, id)
}

func (s *persons) GetRandomWordByLanguageID(ctx context.Context, id string) (string, error) {
	return s.engineAdp.GetRandomWordByLanguageID(ctx, id)
}

func (s *persons) GetRandomOwnNameByLanguageID(ctx context.Context, id string, sex gender.Sex) (string, error) {
	return s.engineAdp.GetRandomOwnNameByLanguageID(ctx, id, sex)
}

func (s *persons) GetRandomCultureNameByLanguageID(ctx context.Context, id string) (string, error) {
	return s.engineAdp.GetRandomCultureNameByLanguageID(ctx, id)
}

func (s *persons) GetRandomReligionNameByLanguageID(ctx context.Context, id string) (string, error) {
	return s.engineAdp.GetRandomReligionNameByLanguageID(ctx, id)
}
