package vwap

import (
	"github.com/geek96/vwap_service/pkg/coinbase"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewVolWeightAvgPrice(t *testing.T)  {
	c := NewVolWeightAvgPrice(10)
	assert.NotNil(t, c)
	assert.Equal(t, 10, c.DataPointsLen)
}

func TestVolWeightedAvgPrice_Process(t *testing.T) {
	c := NewVolWeightAvgPrice(2)
	datach := make(chan coinbase.MatchResponse)
	done := make(chan  bool)
	go c.Process(datach, done)
	datach <- coinbase.MatchResponse{
		Volume: big.NewFloat(0.23),
		Price: big.NewFloat(41519.01),
		ProductId: "BTC-USD",
	}
	done <- true
	p, ok := c.Load("BTC-USD")
	assert.NotNil(t, ok)
	assert.NotNil(t, p)
	prod := p.(VWAPData)
	assert.Equal(t, 1, len(prod.DataPoints))
}

func TestVolWeightedAvgPrice_ProcessInvalidProduct(t *testing.T) {
	c := NewVolWeightAvgPrice(2)
	datach := make(chan coinbase.MatchResponse)
	done := make(chan  bool)
	go c.Process(datach, done)
	datach <- coinbase.MatchResponse{
		Volume: big.NewFloat(0.23),
		Price: big.NewFloat(41519.01),
		ProductId: "BTC-USD",
	}
	done <- true
	p, ok := c.Load("BTC-INVALID")
	assert.NotNil(t, ok)
	assert.Nil(t, p)
}