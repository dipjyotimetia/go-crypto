package store

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (s Store) UpdatePriceInfo(ctx context.Context, symbol, price string) {
	_, err := s.Client.Collection("bnb").Doc("pricing").Set(ctx, map[string]string{
		"symbol": symbol,
		"price":  price,
	})
	if err != nil {
		log.Fatal("error while updating price info", err.Error())
	}
}
