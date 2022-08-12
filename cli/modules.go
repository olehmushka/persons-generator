package cli

import (
	"persons_generator/config"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
)
