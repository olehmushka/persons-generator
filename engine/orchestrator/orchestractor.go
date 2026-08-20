// Package orchestrator is the thin coordination layer between the pure
// generation logic in engine/entities/* and persistence: it either delegates
// straight to a DB-free entity constructor (e.g. CreateCultures, CreateWorld)
// or additionally reads/writes that data through MongoDB (storage) and Redis
// (async world-generation progress). internal/*/adapters/engine wraps this
// package for the ports-and-adapters layer used by handlers/http and cli.
package orchestrator

import (
	"persons_generator/core/mongodb"
	"persons_generator/core/redis"
	"persons_generator/core/wrapped_error"
)

// Orchestrator holds the MongoDB and Redis clients shared by every
// generation/persistence method in this package.
type Orchestrator struct {
	storage redis.Storage
	mongodb mongodb.Client

	dbName string
}

// New connects to MongoDB and Redis using cfg. Both connections are made
// eagerly, so New only succeeds once both backends are reachable.
func New(cfg Config) (*Orchestrator, error) {
	storage, err := redis.NewStorageByParams(cfg.RedisURL, cfg.RedisUsername, cfg.RedisPassword)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create redis storage driver for orch")
	}

	mdb, err := mongodb.New(cfg.MongoDBUsername, cfg.MongoDBPassword, cfg.MongoDBURL, cfg.MongoDBMaxBulkItemsSize)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create mongodb storage driver for orch")
	}

	return &Orchestrator{
		storage: storage,
		mongodb: mdb,
		dbName:  cfg.MongoDBDBName,
	}, nil
}
