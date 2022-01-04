package store

import (
	"context"

	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
)

func (s Store) UpdatePriceInfo(ctx context.Context, symbol, price string) {
	_, _, err := s.Collection("pricing").Add(ctx, model.PriceInfo{
		Symbol: symbol,
		Price:  price,
	})
	if err != nil {
		log.Fatal("error while updating price info", err.Error())
	}
}

func (s Store) PricingHistory(ctx context.Context, priceChange model.PriceChange) {
	_, _, err := s.Collection("pricingHistory").Add(ctx, model.UpdatePriceInfo{
		Symbol:    priceChange.Symbol,
		HighPrice: priceChange.HighPrice,
		LowPrice:  priceChange.LowPrice,
	})
	if err != nil {
		log.Fatal("error while adding price history", err.Error())
	}
}
