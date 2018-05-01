package main

import(
	"fmt";
	"net/http";
	"encoding/json";
	"time"
)

type Message struct{
	text string
	person string
	time string
}
var mess Message

func speak(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(mess)
}

func send(w http.ResponseWriter, r *http.Request){
	if(r.Method == "POST"){
		r.ParseForm()
		mess = Message{text: r.PostFormValue("message"),person: r.PostFormValue("sender"),time: r.PostFormValue("time")}
	}
}

func main() {
	mess = Message{"Welcome to Chat on the Go!","ServerAdmin",time.Now().String()}
	http.HandleFunc("/", speak)
	http.HandleFunc("/send", send)
	fmt.Println("Everything is working and ready...")
	http.ListenAndServe(":8080", nil)
}