#!/bin/bash

echo "Кросскомпиляция на компьютере разработчика auth-proxy:dev"
# env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -tags=jsoniter .
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  .

