package coin

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/go-crypto/internal/store"
)

type CryptoService interface {
	PriceService(ctx context.Context, conn store.CryptoService)
	AveragePriceService(ctx context.Context, conn store.CryptoService)
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
