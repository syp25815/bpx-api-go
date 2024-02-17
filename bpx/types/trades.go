package types

type TradeDetail struct {
	Id            uint64 `json:"id" `
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	QuoteQuantity string `json:"quoteQuantity"`
	Timestamp     int64  `json:"timestamp"`
}
