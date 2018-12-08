#!/bin/bash
source .env
env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build main.go
