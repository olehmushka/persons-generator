package engine

import (
	"context"

	"persons_generator/config"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/culture/entities"

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

func (a *adapter) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*culture.Culture, error) {
	cultures, err := a.engine.CreateCultures(amount, deserializeCulturePreferences(preferred))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create cultures")
	}

	for _, c := range cultures {
		if err := a.engine.SaveCulture(ctx, c); err != nil {
			return nil, err
		}
	}

	return cultures, nil
}

func (a *adapter) GetProtoCultures(ctx context.Context, q string, opts storage.PaginationSortingOpts) ([]*culture.Culture, int, error) {
	out, err := a.engine.FindNativeCultures(ctx, q, opts)
	if err != nil {
		return nil, 0, err
	}
	count, err := a.engine.CountNativeCultures(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return out, count, nil
}

func (a *adapter) GetCultureByID(ctx context.Context, id string) (*culture.Culture, error) {
	return a.engine.ReadCultureByID(ctx, id)
}

func (a *adapter) DeleteCultureByID(ctx context.Context, id string) error {
	return a.engine.DeleteCultureByID(ctx, id)
}

func (a *adapter) DeleteAllCultures(ctx context.Context) error {
	return a.engine.DeleteAllCultures(ctx)
}

func (a *adapter) UpdateCultureLanguage(ctx context.Context, cultureID, languageID string) error {
	return a.engine.UpdateCultureLanguage(ctx, cultureID, languageID)
}
