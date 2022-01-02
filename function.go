package p

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-crypto/router"
	"github.com/rs/cors"
)

func GoCrypto(w http.ResponseWriter, r *http.Request) {
	rc := chi.NewRouter()

	rc.Use(middleware.RealIP)
	rc.Use(middleware.Logger)

	rc.Route("/api/v1", func(rc chi.Router) {
		rc.Get("/trigger", router.Trigger())
		rc.Post("/watchlist", router.Watchlist())
	})

	cors.Default().Handler(rc).ServeHTTP(w, r)
}
