package store

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
)

// https://cloud.google.com/firestore/docs/query-data/queries

type CryptoService interface {
	UpdatePriceInfo(ctx context.Context, coinInfo map[string]string)
	PricingHistory(ctx context.Context, priceChange model.PriceChange)
	AddUpdateWatchlist(ctx context.Context, symbol string, price string)
	DeleteWatchlist(ctx context.Context, symbol string)
	RegisterUser(ctx context.Context, user model.Register)
	LoginUser(ctx context.Context, user model.Login) error
}

type Store struct {
	*firestore.Client
}

// https://pkg.go.dev/cloud.google.com/go/firestore

func NewFireStoreConnection(ctx context.Context) CryptoService {
	client, err := firestore.NewClient(ctx, "dev-aileron-214211") // os.Getenv("PROJECT_ID")
	if err != nil {
		log.Fatal("firestore connection error", err.Error())
	}
	return &Store{client}
}
