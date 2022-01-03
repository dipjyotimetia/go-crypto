package store

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (s Store) UpdatePriceInfo(ctx context.Context, symbol, price string) {
	_, _, err := s.Client.Collection("pricing").Add(ctx, PriceInfo{
		Symbol: symbol,
		Price:  price,
	})
	if err != nil {
		log.Fatal("error while updating price info", err.Error())
	}
}
