#!/bin/sh
GOOS=windows GOARCH=386 go build -o client.exe client.go cliTypes.go
