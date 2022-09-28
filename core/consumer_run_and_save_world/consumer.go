package consumer_run_and_save_world

import (
	"context"
	"persons_generator/config"
	"persons_generator/core/redis"
	"persons_generator/core/wrapped_error"
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
) (RunAndSaveWorldConsumer, error) {
	driver, err := redis.NewConsumer(
		cfg,
		cfg.Redis.Channels.RunAndSaveWorld.Name,
		handler.ProcessRunAndSaveWorld,
		cfg.Redis.Channels.RunAndSaveWorld.ConsumerConcurrency,
	)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create consumer driver for run_and_save_world")
	}
	c := &consumer{
		driver: driver,
	}
	lc.Append(fx.Hook{OnStop: func(context.Context) error {
		return c.driver.GetClient().Close()
	}})

	return c, nil
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
