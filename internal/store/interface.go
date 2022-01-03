package store

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	log "github.com/sirupsen/logrus"
)

type PriceService interface {
	UpdatePriceInfo(ctx context.Context, symbol, price string)
}

type Store struct {
	*firestore.Client
}

func NewFireStoreConnection(ctx context.Context) PriceService {
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal("firestore connection error", err.Error())
	}
	return &Store{client}
}
