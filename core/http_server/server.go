package http_server

import (
	"net/http"

	"persons_generator/config"
	"persons_generator/core/http_server_tools"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type server struct {
	handlers Handlers

	address string
}

func New(cfg *config.Config, handlers Handlers) Server {
	return &server{
		handlers: handlers,

		address: cfg.HTTPServer.Address,
	}
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(Register),
)

func Register(s Server) {
	go s.Register()
}

func (s *server) Register() {
	router := chi.NewRouter()

	router.Use(middleware.Logger, AppendTraceID)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", TraceIDHeader},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           1000,
	}))

	router.Post("/api/cultures", http_server_tools.NewHandlesChain(s.handlers.CreateCultures))
	router.Get("/api/cultures/{id}", http_server_tools.NewHandlesChain(s.handlers.GetCultureByID))
	router.Get("/api/cultures/proto", http_server_tools.NewHandlesChain(s.handlers.GetProtoCultures))
	router.Delete("/api/cultures/{id}", http_server_tools.NewHandlesChain(s.handlers.DeleteCultureByID))
	router.Delete("/api/cultures", http_server_tools.NewHandlesChain(s.handlers.DeleteAllCultures))

	router.Post("/api/religions", http_server_tools.NewHandlesChain(s.handlers.CreateReligions))
	router.Get("/api/religions/{id}", http_server_tools.NewHandlesChain(s.handlers.GetReligionByID))
	router.Delete("/api/religions/{id}", http_server_tools.NewHandlesChain(s.handlers.DeleteReligionByID))
	router.Delete("/api/religions", http_server_tools.NewHandlesChain(s.handlers.DeleteAllReligions))

	router.Post("/api/worlds", http_server_tools.NewHandlesChain(s.handlers.CreateWorld))
	router.Delete("/api/worlds/{id}", http_server_tools.NewHandlesChain(s.handlers.DeleteWorldByID))
	router.Delete("/api/worlds", http_server_tools.NewHandlesChain(s.handlers.DeleteAllWorlds))

	router.Get("/api/persons", http_server_tools.NewHandlesChain(s.handlers.GetWorldPersons))
	router.Delete("/api/persons/{id}", http_server_tools.NewHandlesChain(s.handlers.DeletePersonByID))
	router.Delete("/api/persons", http_server_tools.NewHandlesChain(s.handlers.DeleteAllPersons))
	router.Get("/api/worlds/progress", http_server_tools.NewHandlesChain(s.handlers.GetWorldProgress))

	srv := &http.Server{
		Addr:              s.address,
		Handler:           router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	logrus.WithFields(logrus.Fields{"http_server_address": s.address}).
		Info("starting http server...")
	if err := srv.ListenAndServe(); err != nil {
		logrus.
			Error("can not start http server", err.Error())
	}
}
