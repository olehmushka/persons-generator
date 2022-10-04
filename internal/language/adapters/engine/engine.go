package engine

import (
	"context"
	"fmt"

	"persons_generator/config"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/language/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
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

func (a *adapter) QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	if q == "" {
		return serializeLanguages(a.engine.GetDefaultLanguages(opts)), nil
	}

	langs, err := a.engine.QueryDefaultLanguages(q, opts)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not query default languages (query=%s)", q))
	}

	return serializeLanguages(langs), nil
}

func (a *adapter) CreateLanguage(ctx context.Context, in *entities.Language) (uuid.UUID, error) {
	return a.engine.CreateLanguage(ctx, deserializeLanguage(in))
}

func (a *adapter) ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, error) {
	langs, err := a.engine.ReadLanguagesByName(ctx, name, opts)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read languages by name (name=%s)", name))
	}

	return serializeLanguages(langs), nil
}

func (a *adapter) DeleteLanguageByID(ctx context.Context, id uuid.UUID) error {
	return a.engine.DeleteLanguageByID(ctx, id)
}

func (a *adapter) DeleteAllLanguages(ctx context.Context) error {
	return a.engine.DeleteAllLanguages(ctx)
}
