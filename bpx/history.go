package bpx

import (
	"github.com/spf13/cast"
	"github.com/syp25815/bpx-api-go/bpx/types"
)

//func (c *Client) HistoryOrders(symbol string, limit, offset int64) {
//
//	params := map[string]any{}
//	params["symbol"] = symbol
//	params["limit"] = limit
//	params["offset"] = offset
//	url := API_BASE + "wapi/v1/history/orders"
//	fmt.Println("链接是：", url)
//	_, resp, _ := c.wrapAgent(newAgent().
//		Get(url), params).
//		End()
//
//	fmt.Println(resp)
//	return
//}

func (c *Client) HistoryFills(symbol string, limit, offset int64) (resp []*types.OrderFill) {

	params := map[string]any{}

	if len(symbol) > 0 {
		params["symbol"] = symbol
	}
	params["limit"] = limit
	params["offset"] = offset

	url := API_BASE + "wapi/v1/history/fills"
	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)

	for _, val := range resp {
		val.QuoteQuantity = cast.ToFloat64(val.Price) * cast.ToFloat64(val.Quantity)
		if val.FeeSymbol == "USDC" {
			val.FeeU = cast.ToFloat64(val.Fee)
		} else {
			val.FeeU = cast.ToFloat64(val.Price) * cast.ToFloat64(val.Fee)
		}
	}
	return
}
