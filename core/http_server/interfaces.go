package http_server

import (
	"net/http"
)

type Server interface {
	Register()
}

type Handlers interface {
	CreateCultures(w http.ResponseWriter, r *http.Request)
	GetProtoCultures(w http.ResponseWriter, r *http.Request)
	GetCultureByID(w http.ResponseWriter, r *http.Request)
	DeleteCultureByID(w http.ResponseWriter, r *http.Request)
	DeleteAllCultures(w http.ResponseWriter, r *http.Request)
	UpdateCultureLanguage(w http.ResponseWriter, r *http.Request)

	CreateReligions(w http.ResponseWriter, r *http.Request)
	GetReligionByID(w http.ResponseWriter, r *http.Request)
	DeleteReligionByID(w http.ResponseWriter, r *http.Request)
	DeleteAllReligions(w http.ResponseWriter, r *http.Request)

	GetWorldPersons(w http.ResponseWriter, r *http.Request)
	DeletePersonByID(w http.ResponseWriter, r *http.Request)
	DeleteAllPersons(w http.ResponseWriter, r *http.Request)

	CreateWorld(w http.ResponseWriter, r *http.Request)
	GetWorldProgress(w http.ResponseWriter, r *http.Request)
	GetWorlds(w http.ResponseWriter, r *http.Request)
	DeleteWorldByID(w http.ResponseWriter, r *http.Request)
	DeleteAllWorlds(w http.ResponseWriter, r *http.Request)

	CreateLanguage(w http.ResponseWriter, r *http.Request)
	GetLanguages(w http.ResponseWriter, r *http.Request)
	GetLanguageByID(w http.ResponseWriter, r *http.Request)
	GetDefaultLanguages(w http.ResponseWriter, r *http.Request)
	GetDefaultLanguageSubfamilies(w http.ResponseWriter, r *http.Request)
	DeleteLanguageByID(w http.ResponseWriter, r *http.Request)
	DeleteAllLanguages(w http.ResponseWriter, r *http.Request)
	GetWordLanguageByID(w http.ResponseWriter, r *http.Request)
}
