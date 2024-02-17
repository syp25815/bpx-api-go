package bpx

import (
	"github.com/syp25815/bpx-api-go/bpx/types"
)

func Trades(symbol string, limit uint64) (resp []*types.TradeDetail) {
	params := map[string]any{}
	params["symbol"] = symbol
	params["limit"] = limit
	url := API_BASE + "api/v1/trades"
	newAgent().
		Get(url).Query(params).
		EndStruct(&resp)
	return
}

func HistoryTrades(symbol string, limit, offset int64) (resp []*types.TradeDetail) {
	params := map[string]any{}
	params["symbol"] = symbol
	params["limit"] = limit
	params["offset"] = offset
	url := API_BASE + "api/v1/trades/history"
	newAgent().
		Get(url).
		Query(params).
		EndStruct(&resp)
	return
}
