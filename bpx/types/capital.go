package types

type CapitalDetail struct {
	Available string `json:"available"`
	Locked    string `json:"locked"`
	Staked    string `json:"staked"`
}

type WalletAddress struct {
	Address string `json:"address"`
}

type Deposits struct {
	ConfirmationBlockNumber int     `json:"confirmationBlockNumber"`
	CreatedAt               string  `json:"createdAt"`
	FromAddress             string  `json:"fromAddress"`
	Id                      int     `json:"id"`
	ProviderId              *string `json:"providerId"`
	Quantity                string  `json:"quantity"`
	Source                  string  `json:"source"`
	Status                  string  `json:"status"`
	SubaccountId            *string `json:"subaccountId"`
	Symbol                  string  `json:"symbol"`
	ToAddress               string  `json:"toAddress"`
	TransactionHash         string  `json:"transactionHash"`
}

type WithDrawl struct {
	Blockchain      string      `json:"blockchain"`
	ClientId        string      `json:"clientId"`
	CreatedAt       string      `json:"createdAt"`
	Fee             string      `json:"fee"`
	Id              int         `json:"id"`
	Identifier      string      `json:"identifier"`
	Quantity        string      `json:"quantity"`
	Status          string      `json:"status"`
	SubaccountId    int64       `json:"subaccountId"`
	Symbol          string      `json:"symbol"`
	ToAddress       string      `json:"toAddress"`
	Transaction     interface{} `json:"transaction"`
	TransactionHash string      `json:"transactionHash"`
}
