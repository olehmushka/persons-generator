package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Storage interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, exp time.Duration) error
}

type Consumer interface {
	Consume(context.Context)
	GetClient() *redis.Client
	GetHandler() HandlerFunc
}

type Publisher interface {
	Publish(context.Context, string, []byte) error
}
