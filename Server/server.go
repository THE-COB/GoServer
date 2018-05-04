package main

import(
	"fmt";
	"net/http";
	"encoding/json";
	"time";
)

type Message struct{
	Text string
	Person []string
	Time string
}
var mess Message

func speak(w http.ResponseWriter, r *http.Request){
	b, _ := json.Marshal(mess)
	w.Write(b)
}

func send(w http.ResponseWriter, r *http.Request){
	if(r.Method == "POST"){
		r.ParseForm()
		vals := r.PostForm
		mess = Message{vals["message"][0],vals["sender"],vals["time"][0]}
	}
}

func main() {
	pers := []string{"ServerAdmin", "0"}
	mess = Message{"Welcome to Chat on the Go!",pers,time.Now().String()}
	http.HandleFunc("/", speak)
	http.HandleFunc("/send", send)
	fmt.Println("Everything is working and ready...")
	http.ListenAndServe(":8080", nil)
}