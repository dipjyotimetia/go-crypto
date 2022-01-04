package store

import (
	"context"

	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
)

func (s Store) AddUpdateWatchlist(ctx context.Context, symbol, price string) {
	_, err := s.Client.Collection("watchlist").Doc(symbol).Set(ctx, model.Watch{
		Symbol: symbol,
		Price:  price,
	})
	if err != nil {
		log.Fatal("failed to add watchlist", err.Error())
	}
}

func (s Store) DeleteWatchlist(ctx context.Context, symbol string) {
	_, err := s.Client.Collection("watchlist").Doc(symbol).Delete(ctx)
	if err != nil {
		log.Fatal("failed to delete watchlist", err.Error())
	}
}
