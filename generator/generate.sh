#!/bin/sh

rm -f tlds-alpha-by-domain.txt map.go
curl https://data.iana.org/TLD/tlds-alpha-by-domain.txt | sed '1d' > input.txt
go run convert.go | gofmt > map.go
rm -f input.txt
