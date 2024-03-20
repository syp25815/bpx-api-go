package bpx

import (
	"github.com/syp25815/bpx-api-go/bpx/types"
)

const (
	StatusConfirmed = "confirmed"
)

func (c *Client) Capital() (resp map[string]types.CapitalDetail) {
	url := API_BASE + "api/v1/capital"
	c.wrapAgent(newAgent().
		Get(url), nil).
		EndStruct(&resp)
	return
}

func (c *Client) Deposits(limit, offset int64) (resp []*types.Deposits) {
	params := map[string]any{}
	params["limit"] = limit
	params["offset"] = offset
	url := API_BASE + "wapi/v1/capital/deposits"
	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)
	return
}

func (c *Client) DepositAddress(blockChain string) (resp *types.WalletAddress) {
	params := map[string]any{}
	params["blockchain"] = blockChain
	url := API_BASE + "wapi/v1/capital/deposit/address"
	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)
	return
}

func (c *Client) Withdrawals(limit, offset int64) (resp []*types.WithDrawl) {

	params := map[string]any{}
	params["limit"] = limit
	params["offset"] = offset
	url := API_BASE + "wapi/v1/capital/withdrawals"
	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)
	return
}

/**
* set withdrawal address:
* https://backpack.exchange/settings/withdrawal-addresses?twoFactorWithdrawalAddress=true
 */
func (c *Client) WithdrawalExecute(address, symbol, blockchain, quantity string) (resp *types.WithDrawl) {
	params := map[string]any{}
	params["address"] = address
	params["blockchain"] = blockchain
	params["quantity"] = quantity
	params["symbol"] = symbol

	url := API_BASE + "wapi/v1/capital/withdrawals"
	c.wrapAgent(newAgent().
		Post(url), params).
		EndStruct(&resp)
	return
}
