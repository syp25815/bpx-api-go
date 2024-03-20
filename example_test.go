package bpx_api_go

import (
	"github.com/syp25815/bpx-api-go/bpx"
	"github.com/syp25815/bpx-api-go/xstring"
	"testing"
)

func TestBpx(t *testing.T) {

	var c *bpx.Client

	c = bpx.NewClient("", "")

	//c.Debug = true

	//c.NetInfo()
	//c.Window = "60000"

	for _, val := range c.HistoryOrders("SOL_USDC", "112083497076719617", 100, 0) {
		t.Log(xstring.Json(val))
	}

	//t.Log(xstring.PrettyJson(c.HistoryOrders()))

	for _, val := range c.HistoryFills("", "", 0, 0, 10, 0) {

		t.Log(xstring.Json(val))
	}

	//for _, val := range c.Withdrawals(10, 0) {
	//	t.Log(xstring.Json(val))
	//}

	//t.Log(xstring.PrettyJson(c.WithdrawalExecute("", "USDC", "Solana", "100")))

	//for _, val := range c.OrdersQuery("SOL_USDC") {
	//	t.Log(xstring.Json(val))
	//}

	//t.Log(xstring.Json(c.OrderCancel("SOL_USDC", "111947854823555072", 0)))
	//c.DebugTS = 1708193094831
	//t.Log(xstring.Json(c.OrderQuery("SOL_USDC", "111948072781414400", 0)))

	//t.Log(xstring.Json(c.OrdersQuery("")))

	//t.Log(xstring.Json(c.OrderCancel("SOL_USDC", "111949999842721792", 0)))

	//c.OrdersCancels("SOL_USDC")
	//t.Log(xstring.Json(c.DepositAddress("Solana")))
	//
	//for _, val := range c.Capital() {
	//	t.Log(xstring.Json(val))
	//}
	//
	//for _, val := range c.HistoryFills("SOL_USDC", 10, 0) {
	//	t.Log(xstring.Json(val))
	//}

	//t.Log(xstring.Json(c.OrderExecute("SOL_USDC", "Bid", "Limit", "IOC", "0.1", "111")))

	//c.DebugTS = 1708221111531

	//t.Log(xstring.Json(c.OrderQuery("SOL_USDC", "111949942677897216", 0)))

	//postOnly
	//t.Log(xstring.Json(c.OrderExecute("SOL_USDC", "Bid", "Limit", "", "1", "12")))

}

func TestBpx_pub(t *testing.T) {
	//for _, val := range bpx.Assets() {
	//	t.Log(xstring.Json(val))
	//}
	//
	//for _, val := range bpx.Markets() {
	//	t.Log(xstring.Json(val))
	//}
	//
	//bpx.Ticker("SOL")

	//for _, val := range bpx.Tickers() {
	//	t.Log(xstring.Json(val))
	//}

	//t.Log(xstring.PrettyJson(bpx.Depth("SOL_USDC")))

	//for _, val := range bpx.KLines("SOL_USDC", "1m", 0, 0) {
	//	t.Log(xstring.Json(val))
	//}

	//t.Log(xstring.PrettyJson(bpx.Depth("SOL_USDC")))
	//t.Log(xstring.PrettyJson(bpx.Ticker("SOL_USDC")))

	//t.Log(xstring.Json(bpx.Status()))
	//t.Log(bpx.Ping())
	//t.Log(bpx.SystemTime())
	//
	//for _, val := range bpx.Trades("SOL_USDC", 10) {
	//	t.Log(xstring.Json(val))
	//}
	//for _, val := range bpx.HistoryTrades("SOL_USDC", 10, 0) {
	//	t.Log(xstring.Json(val))
	//}
}
