package store

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	log "github.com/sirupsen/logrus"
)

// https://cloud.google.com/firestore/docs/query-data/queries

type CryptoService interface {
	UpdatePriceInfo(ctx context.Context, symbol, price string)
	AddUpdateWatchlist(ctx context.Context, symbol string, price string)
	DeleteWatchlist(ctx context.Context, symbol string)
}

type Store struct {
	*firestore.Client
}

func NewFireStoreConnection(ctx context.Context) CryptoService {
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal("firestore connection error", err.Error())
	}
	return &Store{client}
}
