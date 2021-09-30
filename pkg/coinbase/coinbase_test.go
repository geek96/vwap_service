package coinbase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cbWebsocketUrl = "wss://ws-feed.pro.coinbase.com"

func TestNewClient(t *testing.T) {
	c, err := NewClient(cbWebsocketUrl)
	defer closeClient(c)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}


func TestNewClientShouldReturnError(t *testing.T) {
	c, err := NewClient("ws://exmaple.com")
	defer closeClient(c)
	assert.Error(t, err)
	assert.Nil(t, c)
}

func TestNewMatchChannelRequest(t *testing.T) {
	expected := []string{"BTC-USD"}
	req := NewMatchChannelRequest(expected)
	assert.NotNil(t, req)
	assert.Equal(t, "matches", req.Name)
	assert.Equal(t, 1, len(req.ProductIds))
	assert.Equal(t, expected, req.ProductIds)
}

func TestCBClient_SubscribeChannel(t *testing.T) {
	c, _ := NewClient(cbWebsocketUrl)
	defer closeClient(c)
	req := []SubscriptionChannel{
		NewMatchChannelRequest([]string{"BTC-USD"}),
	}
	resp , err := c.SubscribeChannel(req)
	assert.NoError(t, err)
	assert.Equal(t, "subscriptions", resp.Type)
}

func TestCBClient_SubscribeChannelShouldReturnError(t *testing.T) {
	c, _ := NewClient(cbWebsocketUrl)
	defer closeClient(c)
	req := []SubscriptionChannel{
		NewMatchChannelRequest([]string{"DUMMY-INVALID"}),
	}
	resp, err := c.SubscribeChannel(req, )
	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "error", resp.Type)
}

func TestCBClient_ReadMatchData(t *testing.T) {
	c , _ := NewClient(cbWebsocketUrl)
	req := []SubscriptionChannel{
		NewMatchChannelRequest([]string{"BTC-USD"}),
	}
	_, _ = c.SubscribeChannel(req)
	respChan := make(chan MatchResponse)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer closeClient(c)
	go c.ReadMatchData(ctx, respChan)
	resp := <- respChan
	assert.NotNil(t, resp)
	assert.Equal(t, "BTC-USD", resp.ProductId)
}

func closeClient(c *CBClient) {
	if c != nil {
		c.Close()
	}
}