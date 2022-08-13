package p

import (
	"context"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/router"
	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

var (
	validate *validator.Validate
	onceFunc sync.Once
)

func GoCrypto(w http.ResponseWriter, r *http.Request) {
	onceBody := func() {
		crypto(w, r)
	}
	go func() {
		onceFunc.Do(onceBody)
	}()
}

func crypto(w http.ResponseWriter, r *http.Request) {
	rc := chi.NewRouter()
	rc.Use(middleware.RealIP)
	rc.Use(middleware.Logger)

	validate = validator.New()
	ctx := context.Background()
	conn := store.NewFireStoreConnection(ctx)

	rc.Route("/api", func(rc chi.Router) {
		rc.Post("/register", router.Register(ctx, validate, conn))
		rc.Post("/login", router.Signin(ctx, validate, conn))
		rc.Post("/reset", router.ResetUser(ctx, validate, conn))
		rc.Post("/logout", router.Logout())
		rc.Get("/refresh", router.Refresh())
		rc.Get("/welcome", router.Welcome())
	})

	rc.Route("/api/v1", func(rc chi.Router) {
		rc.Get("/trigger", router.Trigger(ctx, conn))
		rc.Post("/watchlist", router.AddUpdateWatchlist())
		rc.Delete("/watchlist", router.DeleteWatchlist())
	})

	cors.Default().Handler(rc).ServeHTTP(w, r)
}
