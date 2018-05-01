package main

import(
	"time";
	"bytes";
	"encoding/gob"
)

type Message struct{
	Sender Person
	Text string
	TimeSent time.Time
}

type Person struct{
	Name string
	Id [32]byte
}

func enc(i interface{}) []byte{
	var net bytes.Buffer
	enc := gob.NewEncoder(&net)
	err := enc.Encode(i)
	err = err
	return net.Bytes()
}

func decP(enc []byte) Person{
	var boi Person
	dec := gob.NewDecoder(bytes.NewReader(enc))
	err := dec.Decode(&boi)
	err=err
	boi.Name = boi.Name[:len(boi.Name)-1]
	return boi
}
func decM(enc []byte) Message{
	var mess Message
	dec := gob.NewDecoder(bytes.NewReader(enc))
	err := dec.Decode(&mess)
	err=err
	return mess
}