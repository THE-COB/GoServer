package main

import(
	"time"
)

type Message struct{
	Sender Person
	Text string
	TimeSent time.Time
}

type Person struct{
	Name, Id string
}