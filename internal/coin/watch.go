package coin

import (
	"context"

	"github.com/go-crypto/internal/store"
)

func (b Bnc) AddToWatchList(ctx context.Context, symbol, price string) {
	conn := store.NewFireStoreConnection(ctx)
	conn.AddWatchlist(ctx, symbol, price)
}
