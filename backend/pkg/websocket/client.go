package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int      `json:"type"`
	Body string   `json:"body"`
	Data Response `json:"data"`
}
type Response struct {
	BidValue  float64 `json:"bid_value"`
	Nickname  string  `json:"nickname"`
	UserId    uint    `json:"user_id"`
	AuctionId uint    `json:"auction_id"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var res Response
		_ = json.Unmarshal(message, &res)
		userMessage := Message{Data: Response{
			BidValue:  res.BidValue,
			Nickname:  res.Nickname,
			UserId:    res.UserId,
			AuctionId: res.AuctionId,
		}}
		c.Pool.Broadcast <- userMessage
	}
}
