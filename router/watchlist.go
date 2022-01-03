package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-crypto/internal/coin"
	"github.com/google/uuid"
)

type Watch struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func Watchlist() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		var data *Watch
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		c := coin.NewConnection()
		ref := c.AddToWatchList(ctx, data.Symbol, data.Price)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ref))
	}
}
