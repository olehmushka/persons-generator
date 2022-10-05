package cli

import (
	"persons_generator/config"
	consumerRunAndSaveWorld "persons_generator/core/consumer_run_and_save_world"
	"persons_generator/core/redis"
	mqRunAndSaveWorld "persons_generator/handlers/mq/mq_run_and_save_world"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	cultureServices "persons_generator/internal/culture/services"
	languageEngineAdp "persons_generator/internal/language/adapters/engine"
	languageServices "persons_generator/internal/language/services"
	personsEngineAdp "persons_generator/internal/persons/adapters/engine"
	personsServices "persons_generator/internal/persons/services"
	religionEngineAdp "persons_generator/internal/religion/adapters/engine"
	religionServices "persons_generator/internal/religion/services"
	worldEngineAdp "persons_generator/internal/world/adapters/engine"
	worldMQAdp "persons_generator/internal/world/adapters/mq"
	worldServices "persons_generator/internal/world/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	redis.PublisherModule,

	cultureEngineAdp.Module,
	languageEngineAdp.Module,
	religionEngineAdp.Module,
	personsEngineAdp.Module,
	worldEngineAdp.Module,
	worldMQAdp.Module,

	cultureServices.Module,
	languageServices.Module,
	personsServices.Module,
	religionServices.Module,
	worldServices.Module,

	mqRunAndSaveWorld.Module,
	consumerRunAndSaveWorld.Module,
)
