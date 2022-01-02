package router

import (
	"net/http"

	"github.com/go-crypto/pkg/coin"
	"github.com/google/uuid"
)

func Trigger() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		p := coin.NewConnection()
		p.PriceService()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
