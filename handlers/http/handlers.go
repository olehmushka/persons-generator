package http

import (
	hs "persons_generator/core/http_server"
	cultureServices "persons_generator/internal/culture/services"
	languageServices "persons_generator/internal/language/services"
	personsServices "persons_generator/internal/persons/services"
	religionServices "persons_generator/internal/religion/services"
	worldServices "persons_generator/internal/world/services"

	"go.uber.org/fx"
)

type handlers struct {
	cultureSrv  cultureServices.Culture
	languageSrv languageServices.Language
	personsSrv  personsServices.Persons
	religionSrv religionServices.Religion
	worldSrv    worldServices.World
}

func New(
	cultureSrv cultureServices.Culture,
	languageSrv languageServices.Language,
	personsSrv personsServices.Persons,
	religionSrv religionServices.Religion,
	worldSrv worldServices.World,
) hs.Handlers {
	return &handlers{
		cultureSrv:  cultureSrv,
		languageSrv: languageSrv,
		personsSrv:  personsSrv,
		religionSrv: religionSrv,
		worldSrv:    worldSrv,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)
