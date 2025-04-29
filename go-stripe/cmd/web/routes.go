package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)

	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("./static/*", http.StripPrefix("./static", fileserver))

	return mux
}
