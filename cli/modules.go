package cli

import (
	"persons_generator/config"
	js "persons_generator/core/storage/json_storage"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	cultureServices "persons_generator/internal/culture/services"
	personsEngineAdp "persons_generator/internal/persons/adapters/engine"
	personsServices "persons_generator/internal/persons/services"
	religionEngineAdp "persons_generator/internal/religion/adapters/engine"
	religionServices "persons_generator/internal/religion/services"
	worldEngineAdp "persons_generator/internal/world/adapters/engine"
	worldServices "persons_generator/internal/world/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	js.Module,
	cultureEngineAdp.Module,
	religionEngineAdp.Module,
	personsEngineAdp.Module,
	worldEngineAdp.Module,
	cultureServices.Module,
	personsServices.Module,
	religionServices.Module,
	worldServices.Module,
)
