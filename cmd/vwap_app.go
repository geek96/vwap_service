package main

import (
	"context"
	"flag"
	"github.com/geek96/vwap_service"
	"github.com/geek96/vwap_service/pkg/coinbase"
	"github.com/geek96/vwap_service/pkg/vwap"
	"log"
	"os"
	"os/signal"
)

func main() {
	configFile := flag.String("-c", "config.yaml", "absolute path of the config file")
	flag.Parse()

	conf := vwap_service.NewDefaultConfig()
	conf.Load(*configFile)

	vwp := vwap.NewVolWeightAvgPrice(conf.VWAP.DataPoints)
	client, err := coinbase.NewClient(conf.Coinbase.WebsocketsUrl)
	if err != nil {
		log.Fatalf("connection error %v", err)
	}

	matchReq := []coinbase.SubscriptionChannel{
		coinbase.NewMatchChannelRequest(conf.Coinbase.Channels),
	}
	_, err = client.SubscribeChannel(matchReq)
	if err != nil {
		log.Fatalf("subscription error %v ", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	matchChan := make(chan coinbase.MatchResponse, 20)
	defer close(matchChan)
	done := make(chan bool)
	go vwp.Process(matchChan, done)
	client.ReadMatchData(ctx, matchChan)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Kill)
	select {
		case <- sigs:
			cancel()
	}
	done <- true
}
