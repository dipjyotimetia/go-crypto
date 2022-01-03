package store

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (s Store) AddWatchlist(ctx context.Context, symbol, price string) string {
	ref, _, err := s.Client.Collection("").Add(ctx, Watch{
		Symbol: symbol,
		Price:  price,
	})
	if err != nil {
		log.Fatal("failed to add watchlist", err.Error())
	}
	return ref.ID
}
