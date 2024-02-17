package types

type Asset struct {
	Symbol string `json:"symbol"`
	Tokens []struct {
		Blockchain        string      `json:"blockchain"`
		DepositEnabled    bool        `json:"depositEnabled"`
		MaximumWithdrawal interface{} `json:"maximumWithdrawal"`
		MinimumDeposit    string      `json:"minimumDeposit"`
		MinimumWithdrawal string      `json:"minimumWithdrawal"`
		WithdrawEnabled   bool        `json:"withdrawEnabled"`
		WithdrawalFee     string      `json:"withdrawalFee"`
	} `json:"tokens"`
}

type Ticker struct {
	FirstPrice         string `json:"firstPrice"`
	High               string `json:"high"`
	LastPrice          string `json:"lastPrice"`
	Low                string `json:"low"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	QuoteVolume        string `json:"quoteVolume"`
	Symbol             string `json:"symbol"`
	Trades             int    `json:"trades"`
	Volume             string `json:"volume"`
}

type Depth struct {
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type MarketItem struct {
	BaseSymbol string `json:"baseSymbol"`
	Filters    struct {
		Leverage interface{} `json:"leverage"`
		Price    struct {
			MaxPrice *string `json:"maxPrice"`
			MinPrice string  `json:"minPrice"`
			TickSize string  `json:"tickSize"`
		} `json:"price"`
		Quantity struct {
			MaxQuantity *string `json:"maxQuantity"`
			MinQuantity string  `json:"minQuantity"`
			StepSize    string  `json:"stepSize"`
		} `json:"quantity"`
	} `json:"filters"`
	QuoteSymbol string `json:"quoteSymbol"`
	Symbol      string `json:"symbol"`
}

type KLine struct {
	Close  string `json:"close"`
	End    string `json:"end"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Open   string `json:"open"`
	Start  string `json:"start"`
	Trades string `json:"trades"`
	Volume string `json:"volume"`
}
