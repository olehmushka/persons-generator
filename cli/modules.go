package cli

import (
	"persons_generator/config"
	js "persons_generator/core/storage/json_storage"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	cultureServices "persons_generator/internal/culture/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	js.Module,
	cultureEngineAdp.Module,
	cultureServices.Module,
)
