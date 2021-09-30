package vwap_service

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)


type CoinbaseConf struct {
	WebsocketsUrl string `yaml:"websockets_url"`
	Channels []string `yaml:"channels"`
}

type VWAPConf struct {
	DataPoints int `yaml:"data_points"`
}

type Config struct {
	Coinbase CoinbaseConf `yaml:"coinbase"`
	VWAP     VWAPConf     `yaml:"vwap"`
}


func NewDefaultConfig() *Config {
	return &Config{
		Coinbase: CoinbaseConf{
			WebsocketsUrl: "wss://ws-feed.pro.coinbase.com",
			Channels: []string {
				"BTC-USD", "ETH-USD", "ETH-BTC",
			},
		},
		VWAP: VWAPConf{
			DataPoints: 200,
		},
	}
}

func (c *Config) Load(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("error loading config using the default: [%s]", err.Error())
		return
	}
	if err := yaml.Unmarshal(buf, c); err != nil {
		log.Printf("error unmarshalling yaml using the default: [%s]", err.Error())
	}
}