package coin

import (
	"context"

	"github.com/go-crypto/internal/store"
)

func (b Bnc) AddUpdateWatchList(ctx context.Context, symbol, price string) {
	conn := store.NewFireStoreConnection(ctx)
	conn.AddUpdateWatchlist(ctx, symbol, price)
}

func (b Bnc) DeleteWatchlist(ctx context.Context, symbol string) {
	conn := store.NewFireStoreConnection(ctx)
	conn.DeleteWatchlist(ctx, symbol)
}
