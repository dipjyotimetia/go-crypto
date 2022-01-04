package coin

import (
	"context"
	"encoding/json"

	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/pkg/request"
	log "github.com/sirupsen/logrus"
)

var symbols = []string{"BTCAUD"}

func (b Bnc) PriceService(ctx context.Context) {
	price, err := b.NewListPricesService().Symbol(symbols[0]).Do(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	conn := store.NewFireStoreConnection(ctx)
	priceInfo := map[string]string{}
	for _, symbolPrice := range price {
		priceInfo[symbolPrice.Symbol] = symbolPrice.Price
	}
	conn.UpdatePriceInfo(ctx, symbols[0], priceInfo[symbols[0]])
}

func (b Bnc) AveragePriceService(ctx context.Context) {
	req := request.NewHTTPConn()
	res, err := req.HTTPGet("https://api.binance.com/api/v3/ticker/24hr",
		map[string]string{"Content-Type": "application/json"},
		map[string]string{"symbol": "SHIBAUD"})
	if err != nil {
		log.Fatal(err.Error())
	}
	var priceChange model.PriceChange
	err = json.Unmarshal(res.Body(), &priceChange)
	if err != nil {
		log.Fatal(err.Error())
	}
	conn := store.NewFireStoreConnection(ctx)
	conn.PricingHistory(ctx, priceChange)
}
