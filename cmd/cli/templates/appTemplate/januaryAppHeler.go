package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *application) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}

func (a *application) use(m ...func(http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}

func (a *application) group(s string, m ...func(router chi.Router)) {
	a.App.Routes.Route(s, func(r chi.Router) {
		for _, middleware := range m {
			middleware(r)
		}
	})
}
