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

	CreateReligions(w http.ResponseWriter, r *http.Request)

	GetWorldPersons(w http.ResponseWriter, r *http.Request)

	CreateWorld(w http.ResponseWriter, r *http.Request)
	GetWorldProgress(w http.ResponseWriter, r *http.Request)
}
