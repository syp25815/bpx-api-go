package types

const (
	StatusFilled    = "Filled"
	StatusExpired   = "Expired"
	StatusCancelled = "Cancelled"
)

type Order struct {
	ClientId              uint32 `json:"clientId"`
	CreatedAt             int64  `json:"createdAt"`
	ExecutedQuantity      string `json:"executedQuantity"`
	ExecutedQuoteQuantity string `json:"executedQuoteQuantity"`
	Id                    string `json:"id"`
	OrderType             string `json:"orderType"`
	PostOnly              bool   `json:"postOnly"`
	Price                 string `json:"price"`
	Quantity              string `json:"quantity"`
	SelfTradePrevention   string `json:"selfTradePrevention"`
	Side                  string `json:"side"`
	Status                string `json:"status"`
	Symbol                string `json:"symbol"`
	TimeInForce           string `json:"timeInForce"`
	TriggerPrice          string `json:"triggerPrice"`
}
