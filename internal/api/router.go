package api

import (
	"log/slog"
	"net/http"

	"github.com/example/musician-production-suite/internal/config"
	"github.com/example/musician-production-suite/internal/jobs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Dependencies struct {
	Config  config.Config
	Store   *jobs.Store
	Runner  *jobs.Runner
	Version string
	Logger  *slog.Logger
}

func NewRouter(deps Dependencies) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: deps.Config.CORSAllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}))
	r.Use(requestLogger(deps.Logger))

	handler := NewHandler(deps)
	r.Get("/healthz", handler.Health)
	r.Get("/readyz", handler.Ready)
	r.Handle("/metrics", promhttp.Handler())
	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/jobs", handler.CreateJob)
		api.Get("/jobs/{id}", handler.GetJob)
		api.Get("/jobs/{id}/artifacts/*", handler.GetArtifact)
	})
	return r
}

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			logger.Info("http_request", "method", r.Method, "path", r.URL.Path)
		})
	}
}
