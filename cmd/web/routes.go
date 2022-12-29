package main

import (
	"net/http"

	"github.com/tsawler/bookings/pkg/config"
	"github.com/tsawler/bookings/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// The reason for changing routers
// 1. Clean
// 2. Middleware can process requests and perform actions on it

func routes(app *config.AppConfig) http.Handler {
	// Use the pat library for routes
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
