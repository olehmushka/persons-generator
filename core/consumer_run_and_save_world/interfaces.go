package consumer_run_and_save_world

import "persons_generator/core/redis"

type RunAndSaveWorldConsumer interface {
	GetDriver() redis.Consumer
}
