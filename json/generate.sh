#!/bin/sh

rm -f *.json
go run main.go | jq . > tldmap.json
