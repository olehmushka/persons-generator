package engine

import (
	"context"
	"persons_generator/config"
	"persons_generator/core/storage"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/orchestrator"

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

func (a *adapter) DeletePersonByID(ctx context.Context, id string) error {
	return a.engine.DeletePersonByID(ctx, id)
}

func (a *adapter) DeleteAllPersons(ctx context.Context) error {
	return a.engine.DeleteAllPersons(ctx)
}

func (a *adapter) ReadPersonsByWorldID(ctx context.Context, worldID string, opts storage.PaginationSortingOpts) ([]*person.SerializedPerson, error) {
	return a.engine.ReadPersonsByWorldID(ctx, worldID, opts)
}

func (a *adapter) CountPersonsByWorldID(ctx context.Context, worldID string) (int, error) {
	return a.engine.CountPersonsByWorldID(ctx, worldID)
}
