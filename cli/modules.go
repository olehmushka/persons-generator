package cli

import (
	"persons_generator/config"
	js "persons_generator/core/storage/json_storage"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	cultureServices "persons_generator/internal/culture/services"
	religionEngineAdp "persons_generator/internal/religion/adapters/engine"
	religionServices "persons_generator/internal/religion/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	js.Module,
	cultureEngineAdp.Module,
	religionEngineAdp.Module,
	cultureServices.Module,
	religionServices.Module,
)
