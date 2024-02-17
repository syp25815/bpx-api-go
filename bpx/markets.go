package bpx

import (
	"fmt"
	"github.com/syp25815/bpx-api-go/bpx/types"
)

func Assets() (resp []*types.Asset) {
	newAgent().
		Get(API_BASE + "api/v1/assets").
		EndStruct(&resp)
	return

}

func Markets() (resp []*types.MarketItem) {
	newAgent().
		Get(API_BASE + "api/v1/markets").
		EndStruct(&resp)
	return
}

// 24H交易汇总
func Ticker(symbol string) (resp types.Ticker) {
	url := API_BASE + "api/v1/ticker?symbol=" + symbol
	newAgent().
		Get(url).
		EndStruct(&resp)
	return
}

func Tickers() (resp []*types.Ticker) {
	url := API_BASE + "api/v1/tickers"
	newAgent().
		Get(url).
		EndStruct(&resp)
	return
}

func Depth(symbol string) (resp types.Depth) {
	url := API_BASE + "api/v1/depth?symbol=" + symbol
	newAgent().
		Get(url).
		EndStruct(&resp)
	return
}

func KLines(symbol, interval string, startTime, endTime int64) (resp []*types.KLine) {

	url := API_BASE + "api/v1/klines"

	params := map[string]any{}
	params["symbol"] = symbol
	params["interval"] = interval

	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}

	url = fmt.Sprintf("%s?%s", url, buildQueryParams(params))

	newAgent().
		Get(url).
		EndStruct(&resp)

	return

}
