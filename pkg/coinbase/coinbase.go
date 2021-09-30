package coinbase

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/big"
)

type MatchResponse struct {
	Type string `json:"type"`
	Volume    *big.Float `json:"size,string"`
	Price     *big.Float `json:"price,string"`
	ProductId string     `json:"product_id"`
}

type SubscriptionChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type SubscriptionRequest struct {
	Type     string                `json:"type"`
	Channels []SubscriptionChannel `json:"channels"`
}

type ErrorResponse struct {
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Channels []SubscriptionChannel `json:"channels"`
}

type CBClient struct {
	conn *websocket.Conn
}

//NewClient connect with websocket server and return the
func NewClient(url string) (*CBClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error connecting with coinbase api [%s] ", err.Error()))
	}
	return &CBClient{conn: conn}, nil
}

func NewMatchChannelRequest(productIds []string) SubscriptionChannel {
	return SubscriptionChannel{
		Name:       "matches",
		ProductIds: productIds,
	}
}

func (c *CBClient) Close() {
	if c.conn == nil {
		return
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("error closing connection [%s]", err.Error())
	}
}

func (c *CBClient) SubscribeChannel(channels []SubscriptionChannel) (*ErrorResponse, error) {
	req := SubscriptionRequest{
		Type:     "subscribe",
		Channels: channels,
	}
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, err
	}
	resp := &ErrorResponse{}

	if err := c.conn.ReadJSON(resp); err != nil {
		return nil, err
	}

	if resp.Type == "error" {
		return resp, errors.New(resp.Message)
	}

	if resp.Type == "subscriptions" && len(resp.Channels) == len(channels) {
		log.Println("subscription success")
		return resp, nil
	}
	// received last-mast before the subscription may be a delayed read.
	if resp.Type == "last_match" {
		return resp, nil
	}
	// should not be here
	return nil, errors.New("should not be here")
}

func (c *CBClient) ReadMatchData(ctx context.Context, ch chan MatchResponse) {
	var message MatchResponse
	for {
		select {
		case <-ctx.Done():
			log.Println("finish reading...")
			// closing publish channel!
			close(ch)
			return
		default:
			if err := c.conn.ReadJSON(&message); err != nil {
				log.Printf("error reading match data [%s] \n", err.Error())
				continue
			}

			if message.Type == "match" {
				ch <- message
			}
		}

	}
}
