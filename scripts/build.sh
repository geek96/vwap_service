#!/usr/bin/env bash

echo "Preparing..."

go version

mkdir -p bin

echo "Building..."
go build -o bin/vwap_app ./cmd/vwap_app.go

echo "Copying default configuration..."
cp config.yaml ./bin/

echo "Done!"