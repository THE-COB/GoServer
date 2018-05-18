#!/bin/sh
GOOS=windows GOARCH=386 go build -o server.exe server.go
