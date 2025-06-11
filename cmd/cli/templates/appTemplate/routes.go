package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) defaultRoutes() *chi.Mux {
	// middlewares

	// routes
	a.get("/", a.Handlers.Home)

	// media and static
	fs := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fs))

	return a.App.Routes
}
