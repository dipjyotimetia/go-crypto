package coin

import (
	"context"
	"encoding/json"

	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/pkg/request"
	log "github.com/sirupsen/logrus"
)

const pricingURL = "https://api.binance.com/api/v3/ticker/24hr"

var symbols = []string{
	"BTCAUD",
	"MATICAUD",
	"SHIBAUD",
	"LUNAAUD",
	"ETHAUD",
}

func (b Bnc) PriceService(ctx context.Context, conn store.CryptoService) {
	for _, symbol := range symbols {
		price, err := b.NewListPricesService().Symbol(symbol).Do(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		coinInfo := map[string]string{}
		for _, symbolPrice := range price {
			coinInfo[symbolPrice.Symbol] = symbolPrice.Price
		}
		conn.UpdatePriceInfo(ctx, coinInfo)
	}
}

func (b Bnc) AveragePriceService(ctx context.Context, conn store.CryptoService) {
	req := request.NewHTTPConn()
	for _, symbol := range symbols {
		res, err := req.HTTPGet(pricingURL,
			map[string]string{"Content-Type": "application/json"},
			map[string]string{"symbol": symbol})
		if err != nil {
			log.Fatal(err.Error())
		}
		var priceChange model.PriceChange
		err = json.Unmarshal(res.Body(), &priceChange)
		if err != nil {
			log.Fatal(err.Error())
		}
		conn.PricingHistory(ctx, priceChange)
	}

}
