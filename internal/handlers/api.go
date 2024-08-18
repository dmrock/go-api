package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Rout("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}