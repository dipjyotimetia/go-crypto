package store

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (s Store) AddWatchlist(ctx context.Context, symbol, price string) {
	_, err := s.Client.Collection("watchlist").Doc(symbol).Set(ctx, Watch{
		Symbol: symbol,
		Price:  price,
	})
	if err != nil {
		log.Fatal("failed to add watchlist", err.Error())
	}
}
