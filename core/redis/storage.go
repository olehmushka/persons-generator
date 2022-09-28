package redis

import (
	"context"
	"fmt"
	"persons_generator/config"
	"persons_generator/core/wrapped_error"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

type storage struct {
	client *redis.Client
}

func NewStorage(cfg *config.Config) (Storage, error) {
	return NewStorageByParams(cfg.Redis.URL, cfg.Redis.Username, cfg.Redis.Password)
}

func NewStorageByParams(url, username, password string) (Storage, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "parsing redis url error for storage")
	}
	if username != "" {
		opts.Username = username
	}
	if password != "" {
		opts.Password = password
	}

	client := redis.NewClient(opts)

	return &storage{
		client: client,
	}, nil
}

var StorageModule = fx.Options(
	fx.Provide(NewStorage),
)

func (s *storage) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if redis.Nil.Error() == err.Error() {
			return "", nil
		}
		return "", wrapped_error.NewInternalServerError(err, "can not get value from redis by key")
	}

	return val, nil
}

func (s *storage) Set(ctx context.Context, key, value string, exp time.Duration) error {
	if err := s.client.Set(ctx, key, value, exp).Err(); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not set value in redis (key=%s, value=%s)", key, value))
	}

	return nil
}
