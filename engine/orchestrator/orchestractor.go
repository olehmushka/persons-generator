package orchestrator

import (
	"persons_generator/core/redis"
	"persons_generator/core/wrapped_error"
)

type Orchestrator struct {
	storageFolderName string
	storage           redis.Storage
}

func New(cfg Config) (*Orchestrator, error) {
	storage, err := redis.NewStorageByParams(cfg.RedisURL, cfg.RedisUsername, cfg.RedisPassword)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create redis storage driver for orch")
	}

	return &Orchestrator{
		storageFolderName: cfg.StorageFolderName,
		storage:           storage,
	}, nil
}
