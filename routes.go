package january

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (j *January) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if j.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	j.addRoutes(mux)
	return mux
}

func (j *January) addRoutes(mux *chi.Mux) {
	mux.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "January Running!")
		if err != nil {
			return
		}
	})

}
