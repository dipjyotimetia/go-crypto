package coin

import (
	"context"

	"github.com/go-crypto/internal/store"
	log "github.com/sirupsen/logrus"
)

var symbols = []string{"SHIBAUD"}

func (b Bnc) PriceService(ctx context.Context) {
	price, err := b.NewListPricesService().Symbol(symbols[0]).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	conn := store.NewFireStoreConnection(ctx)
	priceInfo := map[string]string{}
	for _, symbolPrice := range price {
		priceInfo[symbolPrice.Symbol] = symbolPrice.Price
	}
	conn.UpdatePriceInfo(ctx, symbols[0], priceInfo[symbols[0]])
}

func (b Bnc) AveragePriceService(ctx context.Context) {
	price, err := b.NewAveragePriceService().Symbol(symbols[0]).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Info(price)
}
