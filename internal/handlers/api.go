package handlers

import (
	"github.com/AtharvBhujbal/go_API/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorize)
		router.Get("/coins", GetCoinBalance)
	})
}
