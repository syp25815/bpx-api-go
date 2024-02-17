package bpx

import (
	"github.com/syp25815/bpx-api-go/bpx/types"
)

func (c *Client) OrderExecute(symbol, side, orderType, timeInForce, quantity, price string) (resp types.Order) {
	params := map[string]any{}
	params["symbol"] = symbol
	params["side"] = side
	params["orderType"] = orderType
	params["timeInForce"] = timeInForce
	params["quantity"] = quantity
	params["price"] = price

	url := API_BASE + "api/v1/order"
	c.wrapAgent(newAgent().
		Post(url), params).
		EndStruct(&resp)
	return

}
