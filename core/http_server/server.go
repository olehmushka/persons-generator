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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Trace-ID"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           1000,
	}))

	router.Post("/api/culture", http_server_tools.NewHandlesChain(s.handlers.CreateCultures))

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
