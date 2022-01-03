package coin

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	API_KEY    = os.Getenv("BNC_API_KEY")
	SECRET_KEY = os.Getenv("BNC_SECRET_KEY")
)

type CryptoService interface {
	PriceService(ctx context.Context)
	AveragePriceService(ctx context.Context)
	AddToWatchList(ctx context.Context, symbol string, price string)
}

type Bnc struct {
	*binance.Client
}

func NewConnection() CryptoService {
	client := binance.NewClient(API_KEY, SECRET_KEY)
	return &Bnc{client}
}
