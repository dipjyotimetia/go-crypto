package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-crypto/internal/coin"
	"github.com/google/uuid"
)

func AddUpdateWatchlist() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		type Watch struct {
			Symbol string `json:"symbol"`
			Price  string `json:"price"`
		}
		var data *Watch
		ctx := context.Background()
		c := coin.NewConnection()
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c.AddUpdateWatchList(ctx, data.Symbol, data.Price)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s added to watchlist", data.Symbol)))
	}
}

func DeleteWatchlist() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		type Watch struct {
			Symbol string `json:"symbol"`
		}
		var data *Watch
		ctx := context.Background()
		c := coin.NewConnection()
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c.DeleteWatchlist(ctx, data.Symbol)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s deleted from watchlist", data.Symbol)))
	}
}
