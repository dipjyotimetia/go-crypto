package model

import (
	"time"
)

type PriceChange struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type UpdatePriceInfo struct {
	Symbol         string    `firestore:"symbol,omitempty"`
	HighPrice      string    `firestore:"highPrice,omitempty"`
	LowPrice       string    `firestore:"lowPrice,omitempty"`
	OccurrenceTime time.Time `firestore:"occurrenceTime"`
}

type PriceInfo struct {
	Symbol         string    `firestore:"symbol,omitempty"`
	Price          string    `firestore:"price,omitempty"`
	OccurrenceTime time.Time `firestore:"occurrenceTime"`
}

type Watch struct {
	Symbol         string    `firestore:"symbol,omitempty"`
	Price          string    `firestore:"price,omitempty"`
	OccurrenceTime time.Time `firestore:"occurrenceTime"`
}
