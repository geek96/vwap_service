# VWAP (Volume Weight Average Price)

This service use the coinbase websocket match api to calculate the VWAP of given pairs.

## Configure application
Update the following configuration in the config.yaml file at root dir.
```yaml
coinbase:
  websocket_url: wss://ws-feed.pro.coinbase.com
  channels:
    - BTC-USD
    - ETH-USD
    - ETH-BTC
vwap:
  datapoints: 200
```


## Download dependencies
```bash
go mod download
```

## Tests
```bash
go test ./...
```

## Build
```bash
./scripts/build.sh
```
If you want to change the update the config after building the app config.yaml is also copied in the bin dir `bin/config.yaml`, by default it loads the
config file relative to executable.


## Run
```bash
./bin/vwap_app
```
If you want to load the configuration file from any other path, use `-c=path-of-config-path`.

## Dependencies
* Gorilla Websocket 
* Stretchr Testify

## Code Structure
```
.
├── README.md
├── cmd
│   └── vwap_app.go
├── config.go
├── config.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── coinbase
│   │   ├── coinbase.go
│   │   └── coinbase_test.go
│   └── vwap
│       ├── vwap.go
│       ├── vwap_stream.go
│       └── vwap_test.go
└── scripts
    └── build.sh
```
