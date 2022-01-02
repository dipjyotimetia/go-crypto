package router

import (
	"net/http"

	"github.com/google/uuid"
)

func Watchlist() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.New().String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
