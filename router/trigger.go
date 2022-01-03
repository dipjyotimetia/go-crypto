package router

import (
	"context"
	"net/http"

	"github.com/go-crypto/internal/coin"
	"github.com/google/uuid"
)

func Trigger() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		ctx := context.Background()
		p := coin.NewConnection()
		p.PriceService(ctx)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
