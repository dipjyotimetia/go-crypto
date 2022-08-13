package store

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
)

// https://cloud.google.com/firestore/docs/query-data/queries

type Service interface {
	CryptoService
	UserService
}

type CryptoService interface {
	UpdatePriceInfo(ctx context.Context, coinInfo map[string]string)
	PricingHistory(ctx context.Context, priceChange model.PriceChange)
	AddUpdateWatchlist(ctx context.Context, symbol string, price string)
	DeleteWatchlist(ctx context.Context, symbol string)
}

type UserService interface {
	RegisterUser(ctx context.Context, user model.Register) error
	LoginUser(ctx context.Context, user model.Login) error
	ResetPassword(ctx context.Context, reset model.ResetPassword) error
}

type Store struct {
	*firestore.Client
}

// https://pkg.go.dev/cloud.google.com/go/firestore
// https://github.com/GoogleCloudPlatform/golang-samples/tree/main/firestore

func NewFireStoreConnection(ctx context.Context) Service {
	client, err := firestore.NewClient(ctx, "dev-aileron-214211") // TODO: os.Getenv("PROJECT_ID")
	if err != nil {
		log.Fatal("firestore connection error", err.Error())
	}
	return &Store{client}
}
