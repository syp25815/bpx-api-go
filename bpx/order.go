package bpx

import (
	"github.com/syp25815/bpx-api-go/bpx/types"
)

func (c *Client) OrderExecute(symbol, side, orderType, timeInForce, quantity, price string) (resp *types.Order) {
	params := map[string]any{}
	params["symbol"] = symbol
	params["side"] = side
	params["orderType"] = orderType
	if len(timeInForce) < 1 {
		params["postOnly"] = true
	} else {
		params["timeInForce"] = timeInForce
	}
	params["quantity"] = quantity
	params["price"] = price

	url := API_BASE + "api/v1/order"
	c.wrapAgent(newAgent().
		Post(url), params).
		EndStruct(&resp)
	return

}

//func (c *Client) OrderQuery(symbol, orderId string, clientId uint32) (resp *types.Order) {
//	params := map[string]any{}
//	params["symbol"] = symbol
//	params["orderId"] = orderId
//	if clientId > 0 {
//		params["clientId"] = clientId
//	}
//	url := API_BASE + "api/v1/order"
//
//	_, dd, err := c.wrapAgent(newAgent().
//		Get(url), params).
//		End()
//	fmt.Println(err)
//	fmt.Println(dd)
//	return
//
//}

func (c *Client) OrdersQuery(symbol string) (resp []*types.Order) {
	params := map[string]any{}

	if len(symbol) > 0 {
		params["symbol"] = symbol
	}

	url := API_BASE + "api/v1/orders"

	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)
	return

}

func (c *Client) OrderCancel(symbol, orderId string, clientId uint32) (resp *types.Order) {
	params := map[string]any{}
	params["symbol"] = symbol
	params["orderId"] = orderId
	if clientId > 0 {
		params["clientId"] = clientId
	}
	url := API_BASE + "api/v1/order"
	c.wrapAgent(newAgent().Delete(url),
		params).
		EndStruct(&resp)

	return
}

func (c *Client) OrdersCancels(symbol string) (resp []*types.Order) {
	params := map[string]any{}
	params["symbol"] = symbol
	url := API_BASE + "api/v1/orders"
	c.wrapAgent(newAgent().
		Delete(url), params).
		EndStruct(&resp)
	return

}
