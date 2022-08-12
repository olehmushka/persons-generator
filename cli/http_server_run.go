package cli

import (
	hs "persons_generator/core/http_server"
	"persons_generator/core/zerodowntime"
	httpHandlers "persons_generator/handlers/http"

	"go.uber.org/fx"
)

const HTTPServerRunCommand = "http_server_run"

func runHTTPServerCommand() error {
	app := fx.New(
		fx.StartTimeout(defaultAppStartTimeout),
		fx.StopTimeout(defaultAppStopTimeout),
		fx.Options(
			Modules,

			httpHandlers.Module,
			hs.Module,
		),
	)

	return zerodowntime.HandleApp(app)
}
