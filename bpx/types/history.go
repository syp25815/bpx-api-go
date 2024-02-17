package types

type OrderFill struct {
	Fee           string  `json:"fee"`
	FeeSymbol     string  `json:"feeSymbol"`
	FeeU          float64 `json:"feeU"`
	Id            int     `json:"id"`
	IsMaker       bool    `json:"isMaker"`
	OrderId       string  `json:"orderId"`
	Price         string  `json:"price"`
	Quantity      string  `json:"quantity"`
	Side          string  `json:"side"`
	Symbol        string  `json:"symbol"`
	Timestamp     string  `json:"timestamp"`
	TradeId       int     `json:"tradeId"`
	QuoteQuantity float64 `json:"quote_quantity"`
}
