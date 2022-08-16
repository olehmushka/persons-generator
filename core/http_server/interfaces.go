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
}
