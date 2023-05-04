#!/bin/sh
set -xe

rm -rf release && mkdir -p release
env CGO_ENABLED=0 GO111MODULE=on go test -v -timeout 30s -covermode=count -coverprofile=./release/coverage.out -coverpkg ./...

go tool cover -html=./release/coverage.out -o ./release/coverage.html
go tool cover -func=./release/coverage.out -o ./release/coverage.txt
tail -n 1 ./release/coverage.txt | awk '{print $1,$3}'