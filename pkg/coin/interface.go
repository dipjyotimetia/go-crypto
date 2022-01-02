package coin

import (
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	API_KEY    = os.Getenv("BNC_API_KEY")
	SECRET_KEY = os.Getenv("BNC_SECRET_KEY")
)

type CryptoService interface {
	PriceService()
	AveragePriceService()
}

type Bnc struct {
	*binance.Client
}

func NewConnection() CryptoService {
	client := binance.NewClient(API_KEY, SECRET_KEY)
	return &Bnc{client}
}
