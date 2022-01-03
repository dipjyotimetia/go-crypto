package coin

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
)

type CryptoService interface {
	PriceService(ctx context.Context)
	AveragePriceService(ctx context.Context)
	AddUpdateWatchList(ctx context.Context, symbol string, price string)
	DeleteWatchlist(ctx context.Context, symbol string)
}

type Bnc struct {
	*binance.Client
}

func NewConnection() CryptoService {
	client := binance.NewClient(os.Getenv("BNC_API_KEY"), os.Getenv("BNC_SECRET_KEY"))
	return &Bnc{client}
}
