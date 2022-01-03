package coin

import (
	"context"

	"github.com/go-crypto/internal/store"
)

func (b Bnc) AddToWatchList(ctx context.Context, symbol, price string) string {
	conn := store.NewFireStoreConnection(ctx)
	return conn.AddWatchlist(ctx, symbol, price)
}
