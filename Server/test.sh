#!/bin/sh
gnome-terminal -x go run ./server.go
go run ../Client/client.go
