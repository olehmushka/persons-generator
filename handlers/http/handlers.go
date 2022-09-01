package http

import (
	hs "persons_generator/core/http_server"
	cultureServices "persons_generator/internal/culture/services"
	religionServices "persons_generator/internal/religion/services"

	"go.uber.org/fx"
)

type handlers struct {
	cultureSrv  cultureServices.Culture
	religionSrv religionServices.Religion
}

func New(cultureSrv cultureServices.Culture, religionSrv religionServices.Religion) hs.Handlers {
	return &handlers{
		cultureSrv:  cultureSrv,
		religionSrv: religionSrv,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)
