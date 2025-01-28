package router

import (
	"github.com/alaa-aqeel/todo/src/api/http/handlers/user_handler"
	"github.com/go-chi/chi/v5"
)

func NewChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			userHandler := user_handler.New()

			r.Post("/", userHandler.Store)
		})
	})
	return r
}
