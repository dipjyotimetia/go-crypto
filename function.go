package p

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/router"
	"github.com/rs/cors"
)

func GoCrypto(w http.ResponseWriter, r *http.Request) {
	rc := chi.NewRouter()
	rc.Use(middleware.RealIP)
	rc.Use(middleware.Logger)

	ctx := context.Background()
	conn := store.NewFireStoreConnection(ctx)

	rc.Route("/api/v1", func(rc chi.Router) {
		rc.Get("/trigger", router.Trigger(ctx, conn))
		rc.Post("/watchlist", router.AddUpdateWatchlist())
		rc.Delete("/watchlist", router.DeleteWatchlist())
	})

	cors.Default().Handler(rc).ServeHTTP(w, r)
}
