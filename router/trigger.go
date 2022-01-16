package router

import (
	"context"
	"net/http"

	"github.com/go-crypto/internal/coin"
	"github.com/go-crypto/internal/store"
	"github.com/google/uuid"
)

func Trigger(ctx context.Context, conn store.CryptoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		p := coin.NewConnection()
		p.PriceService(ctx, conn)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
