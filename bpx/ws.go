package bpx

import (
	"github.com/spf13/cast"
	"time"
)

func (c *Client) WsOrderUpdate() map[string]any {
	ms := cast.ToString(time.Now().UnixMilli())
	params := map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{"account.orderUpdate"},
		"signature": []string{
			c.Key,
			c.sign("subscribe", ms, nil),
			ms,
			c.Window,
		},
	}
	return params
}
