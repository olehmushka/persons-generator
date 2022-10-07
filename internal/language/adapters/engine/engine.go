package engine

import (
	"context"
	"fmt"

	"persons_generator/config"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/language/entities"

	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
		RedisURL:      cfg.Redis.URL,
		RedisUsername: cfg.Redis.Username,
		RedisPassword: cfg.Redis.Password,

		MongoDBURL:              cfg.MongoDB.URL,
		MongoDBUsername:         cfg.MongoDB.Username,
		MongoDBPassword:         cfg.MongoDB.Password,
		MongoDBMaxBulkItemsSize: cfg.MongoDB.MaxBulkItemsSize,
		MongoDBDBName:           cfg.MongoDB.DBName,
	})
	if err != nil {
		return nil, err
	}

	return &adapter{engine: e}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) QueryDefaultLanguages(ctx context.Context, q string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	if q == "" {
		return serializeLanguages(a.engine.GetDefaultLanguages(ctx, opts)), nil
	}

	langs, err := a.engine.QueryDefaultLanguages(ctx, q, opts)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not query default languages (query=%s)", q))
	}

	return serializeLanguages(langs), nil
}

func (a *adapter) CountDefaultLanguages(ctx context.Context, q string) (int, error) {
	return a.engine.CountDefaultLanguages(ctx, q)
}

func (a *adapter) CreateLanguage(ctx context.Context, in *entities.Language) (string, error) {
	return a.engine.CreateLanguage(ctx, deserializeLanguage(in))
}

func (a *adapter) ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	langs, err := a.engine.ReadLanguagesByName(ctx, name, opts)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read languages by name (name=%s)", name))
	}

	return serializeLanguages(langs), nil
}

func (a *adapter) CountLanguagesByName(ctx context.Context, name string) (int, error) {
	return a.engine.CountLanguagesByName(ctx, name)
}

func (a *adapter) ReadDefaultLanguageSubfamilies(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.Subfamily, error) {
	return serializeSubfamilies(a.engine.GetDefaultLanguageSubfamilies(opts)), nil
}

func (a *adapter) CountDefaultLanguageSubfamilies(ctx context.Context) (int, error) {
	return a.engine.CountDefaultLanguageSubfamilies()
}

func (a *adapter) DeleteLanguageByID(ctx context.Context, id string) error {
	return a.engine.DeleteLanguageByID(ctx, id)
}

func (a *adapter) DeleteAllLanguages(ctx context.Context) error {
	return a.engine.DeleteAllLanguages(ctx)
}

func (a *adapter) ReadLanguageByID(ctx context.Context, id string) (*entities.Language, error) {
	lang, err := a.engine.ReadLanguageByID(ctx, id)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not read language by id")
	}

	return serializeLanguage(lang), nil
}

func (a *adapter) GetRandomWordByLanguageID(ctx context.Context, id string) (string, error) {
	return a.engine.GetRandomWordByLanguageID(ctx, id)
}

func (a *adapter) GetRandomOwnNameByLanguageID(ctx context.Context, id string, sex gender.Sex) (string, error) {
	return a.engine.GetRandomOwnNameByLanguageID(ctx, id, sex)
}

func (a *adapter) GetRandomCultureNameByLanguageID(ctx context.Context, id string) (string, error) {
	return a.engine.GetRandomCultureNameByLanguageID(ctx, id)
}

func (a *adapter) GetRandomReligionNameByLanguageID(ctx context.Context, id string) (string, error) {
	return a.engine.GetRandomReligionNameByLanguageID(ctx, id)
}
