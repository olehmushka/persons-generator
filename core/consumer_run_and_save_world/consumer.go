package consumer_run_and_save_world

import (
	"context"
	"persons_generator/config"
	"persons_generator/core/redis"
	mqRunAndSaveWorld "persons_generator/handlers/mq/mq_run_and_save_world"

	"go.uber.org/fx"
)

type consumer struct {
	driver redis.Consumer
}

func New(
	lc fx.Lifecycle,
	cfg *config.Config,
	handler mqRunAndSaveWorld.RunAndSaveWorld,
) RunAndSaveWorldConsumer {
	c := &consumer{
		driver: redis.NewConsumer(
			cfg,
			cfg.Redis.Channels.RunAndSaveWorld.Name,
			handler.ProcessRunAndSaveWorld,
			cfg.Redis.Channels.RunAndSaveWorld.ConsumerConcurrency,
		),
	}
	lc.Append(fx.Hook{OnStop: func(context.Context) error {
		return c.driver.GetClient().Close()
	}})

	return c
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(consumer RunAndSaveWorldConsumer) {
		go consumer.GetDriver().Consume(context.Background())
	}),
)

func (c *consumer) GetDriver() redis.Consumer {
	return c.driver
}
