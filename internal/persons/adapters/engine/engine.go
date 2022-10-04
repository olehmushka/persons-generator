package engine

import (
	"context"
	"persons_generator/config"
	"persons_generator/engine/orchestrator"

	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
		StorageFolderName: cfg.JSONStorage.StorageFolder,
		RedisURL:          cfg.Redis.URL,
		RedisUsername:     cfg.Redis.Username,
		RedisPassword:     cfg.Redis.Password,

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

func (a *adapter) DeleteAllPersons(ctx context.Context) error {
	return a.engine.DeleteAllPersons(ctx)
}
