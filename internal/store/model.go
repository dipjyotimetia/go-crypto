package store

type PriceInfo struct {
	Symbol string `firestore:"symbol,omitempty"`
	Price  string `firestore:"price,omitempty"`
}

type Watch struct {
	Symbol string `firestore:"symbol,omitempty"`
	Price  string `firestore:"price,omitempty"`
}
