package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Contetent-Type", "X-CSRF-Token", "X-Requested-With"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Get("/api/payment-intent", app.GetPyamentIntent)

	return mux
}
