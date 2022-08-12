package cli

import (
	"persons_generator/config"
	cultureEngineAdp "persons_generator/internal/culture/adapters/engine"
	cultureServices "persons_generator/internal/culture/services"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	cultureEngineAdp.Module,
	cultureServices.Module,
)
