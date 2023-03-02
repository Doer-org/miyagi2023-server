package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Doer-org/miyagi2023-server/pkg/adapter/http/handler"
	"github.com/Doer-org/miyagi2023-server/pkg/adapter/http/middleware"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/registry"
)

func New(repository *registry.Repository) http.Handler {

	r := chi.NewRouter()

	middleware.SetCommonMiddlewares(r)

	userHandler := handler.NewUser(repository)
	spotHandler := handler.NewSpot(repository)
	stampCardHandler := handler.NewStampCard(repository)
	stampLogHandler := handler.NewStampLog(repository)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.Get)
		r.Post("/", userHandler.Create)
	})

	r.Route("/spots", func(r chi.Router) {
		r.Get("/{id}", spotHandler.Get)
		r.Post("/", spotHandler.Create)
		r.Get("/list", spotHandler.List)
	})

	r.Route("/stamp_cards", func(r chi.Router) {
		r.Get("/{id}", stampCardHandler.Get)
		r.Post("/", stampCardHandler.Create)
		r.Get("/list", stampCardHandler.List)
	})

	r.Route("/stamp_logs", func(r chi.Router) {
		r.Post("/", stampLogHandler.Create)
		r.Get("/list", stampLogHandler.List)
	})

	return r
}
