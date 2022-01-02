package coin

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var symbols = []string{"SHIBAUD"}

func (b Bnc) PriceService() {
	price, err := b.NewListPricesService().Symbol(symbols[0]).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}

func (b Bnc) AveragePriceService() {
	price, err := b.NewAveragePriceService().Symbol(symbols[0]).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
