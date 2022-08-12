package http

import (
	hs "persons_generator/core/http_server"

	"go.uber.org/fx"
)

type handlers struct{}

func New() hs.Handlers {
	return &handlers{}
}

var Module = fx.Options(
	fx.Provide(New),
)
