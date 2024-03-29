package bpx

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cast"
	"github.com/syp25815/bpx-api-go/bpx/types"
	"sort"
	"strings"
	"time"
)

const (
	API_BASE = "https://api.backpack.exchange/"
	WSS_BASE = "wss://ws.backpack.exchange/"
)

const (
	VERSION  = "1.0.3"
	TIME_OUT = time.Second * 6
)

type Client struct {
	Key     string
	Secret  string
	Proxy   string
	Window  string
	DebugTS int64
	Debug   bool
}

func NewClient(key, secret string) *Client {
	return &Client{
		Key:    key,
		Secret: secret,
		Window: "6000",
	}
}

func newAgent() *gorequest.SuperAgent {
	return gorequest.
		New().
		Timeout(TIME_OUT)
}

func (c *Client) wrapAgent(request *gorequest.SuperAgent, params map[string]any) *gorequest.SuperAgent {
	ms := cast.ToString(time.Now().UnixNano() / 1e6)
	if request.Method == gorequest.GET {
		request = request.Query(buildQueryParams(params))
	} else {
		request = request.SendMap(params)
	}

	instruction := ""

	url := request.Url
	method := request.Method

	if c.Debug {
		fmt.Println("request url：", url)
	}
	if strings.HasSuffix(url, "/api/v1/capital") {
		if method == gorequest.GET {
			instruction = "balanceQuery"
		}
	} else if strings.HasSuffix(url, "/wapi/v1/capital/deposits") {
		if method == gorequest.GET {
			instruction = "depositQueryAll"
		}
	} else if strings.HasSuffix(url, "/wapi/v1/history/fills") {
		if method == gorequest.GET {
			instruction = "fillHistoryQueryAll"
		}
	} else if strings.HasSuffix(url, "/wapi/v1/history/orders") {
		if method == gorequest.GET {
			instruction = "orderHistoryQueryAll"
		}
	} else if strings.HasSuffix(url, "/wapi/v1/capital/deposit/address") {
		if method == gorequest.GET {
			instruction = "depositAddressQuery"
		}
	} else if strings.HasSuffix(url, "/wapi/v1/capital/withdrawals") {
		if method == gorequest.GET {
			instruction = "withdrawalQueryAll"
		} else if method == gorequest.POST {
			instruction = "withdraw"
		}
	} else if strings.HasSuffix(url, "/api/v1/order") {
		if method == gorequest.GET {
			instruction = "orderQuery"
		} else if method == gorequest.POST {
			instruction = "orderExecute"
		} else if method == gorequest.DELETE {
			instruction = "orderCancel"
		}
	} else if strings.HasSuffix(url, "/api/v1/orders") {
		if method == gorequest.GET {
			instruction = "orderQueryAll"
		} else if method == gorequest.DELETE {
			instruction = "orderCancelAll"
		}
	}

	return request.
		Set("Content-Type", "application/json; charset=utf-8").
		Set("X-Window", c.Window).
		Set("X-API-Key", c.Key).
		Set("X-Timestamp", ms).
		Set("X-Signature", c.sign(instruction, ms, params)).
		SetDebug(c.Debug).
		Proxy(c.Proxy)
}

func (c *Client) sign(instruction, ms string, params map[string]any) string {
	signStrBuilder := strings.Builder{}

	if len(instruction) > 0 {
		signStrBuilder.WriteString(fmt.Sprintf("instruction=%s", instruction))
	}

	paramsStr := buildQueryParams(params)

	if len(paramsStr) > 0 {
		signStrBuilder.WriteString("&" + paramsStr)
	}

	if c.Debug && c.DebugTS > 0 {
		ms = cast.ToString(c.DebugTS)
	}

	signStrBuilder.WriteString(fmt.Sprintf("&timestamp=%s&window=%s", ms, c.Window))

	if c.Debug {
		fmt.Println("Signature Str:", signStrBuilder.String())
	}

	apiSecret, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(c.Secret))

	pki := ed25519.NewKeyFromSeed(apiSecret)
	return base64.StdEncoding.EncodeToString(ed25519.Sign(pki, []byte(signStrBuilder.String())))
}

func buildQueryParams(val map[string]any) string {
	p := strings.Builder{}
	keys := make([]string, 0, len(val))
	for k, _ := range val {
		if len(k) > 0 {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		p.WriteString(fmt.Sprintf(`%s=%v&`, k, val[k]))
	}
	tempStr := p.String()
	if strings.HasSuffix(tempStr, "&") &&
		len(tempStr) > 0 {
		tempStr = tempStr[:len(tempStr)-1]
	}

	return tempStr
}

func (c *Client) NetInfo() (resp *types.NetInfo) {
	url := `https://ipinfo.io`
	c.wrapAgent(newAgent().
		Get(url),
		nil).
		EndStruct(&resp)
	return
}
