package bpx_api_go

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/syp25815/bpx-api-go/bpx"
	"github.com/syp25815/bpx-api-go/bpx/types"
	"github.com/syp25815/bpx-api-go/xstring"
	"log"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"
)

func TestSocketMsg(t *testing.T) {

	//bpx := bpx.NewClient("", "")
	flag.Parsed()

	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial(bpx.WSS_BASE, nil)

	t.Log(err)

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	//c.WriteJSON(bpx.WsOrderUpdate())

	//c.WriteJSON(map[string]any{
	//	"method": "SUBSCRIBE",
	//	"params": []string{"depth.SOL_USDC"},
	//})
	//
	//c.WriteJSON(map[string]any{
	//	"method": "SUBSCRIBE",
	//	"params": []string{"trade.SOL_USDC"},
	//})
	//
	//c.WriteJSON(map[string]any{
	//	"method": "SUBSCRIBE",
	//	"params": []string{"bookTicker.SOL_USDC"},
	//})
	//c.WriteJSON(map[string]any{
	//	"method": "SUBSCRIBE",
	//	"params": []string{"kline.1m.SOL_USDC"},
	//})
	c.WriteJSON(map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{"ticker.SOL_USDC", "depth.SOL_USDC"},
	})

	go func() {
		defer close(done)
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s, type: %s", message, websocket.FormatMessageType(mt))

			var wsResp types.WsResp
			xstring.OmitDefaultAPI.Unmarshal(message, &wsResp)
			t.Log(wsResp.Stream)
			dataSpit := strings.Split(wsResp.Stream, ".")

			if len(dataSpit) > 1 {
				switch dataSpit[0] {
				case "kline":
				case "ticker":
				case "trade":
				case "depth":
				case "bookTicker":

				}
			}
			log.Printf("recv: %s, type: %d", message, mt)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			log.Println("来了？？", t.Unix())
			c.WriteMessage(websocket.TextMessage, []byte("Pong"))
			//err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
