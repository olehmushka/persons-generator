package config

type Redis struct {
	Addr     string `env:"REDIS_ADDR"`
	Channels struct {
		RunAndSaveWorld struct {
			Name                string `default:"run_and_save_world" env:"REDIS_CH_NAME_RUN_AND_SAVE_WORLD"`
			ConsumerConcurrency int    `default:"1" env:"REDIS_CH_CONCURRENCY_RUN_AND_SAVE_WORLD"`
		}
	}
}
