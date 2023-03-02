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

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.Get)
		r.Post("/", userHandler.Create)
	})

	return r
}
