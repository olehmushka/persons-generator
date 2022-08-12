package http

import (
	hs "persons_generator/core/http_server"
	cultureServices "persons_generator/internal/culture/services"

	"go.uber.org/fx"
)

type handlers struct {
	cultureSrv cultureServices.Culture
}

func New(cultureSrv cultureServices.Culture) hs.Handlers {
	return &handlers{
		cultureSrv: cultureSrv,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)
