package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"persons_generator/config"
	c "persons_generator/core/context"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

type publisher struct {
	client *redis.Client
}

func NewPublisher(cfg *config.Config) Publisher {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})

	return &publisher{
		client: client,
	}
}

var PublisherModule = fx.Options(
	fx.Provide(NewPublisher),
)

func (p *publisher) Publish(ctx context.Context, ch string, v []byte) error {
	msg := Message{
		Data:      v,
		Timestamp: tools.SerializeTime(time.Now()),
		TraceID:   c.GetTraceID(ctx),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not marshal redis message")
	}

	if err := p.client.Publish(ctx, ch, b).Err(); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not publish redis msg (channel=%s, msg=%+v)", ch, msg))
	}

	return nil
}
