package cli

import (
	"persons_generator/config"
	js "persons_generator/core/storage/json_storage"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	jsAdp "persons_generator/internal/culture/adapters/json_storage"
	cultureServices "persons_generator/internal/culture/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	js.Module,
	cultureEngineAdp.Module,
	jsAdp.Module,
	cultureServices.Module,
)
