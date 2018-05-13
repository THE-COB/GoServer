package main

import(
	"fmt";
	"net/http";
	"encoding/json";
	"time";
)

type Message struct{
	Text string
	Name string
	Id [32]byte
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
		var pId [32]byte
		copy(pId[:], vals["sender"][1])
		mess = Message{vals["message"][0],vals["sender"][0],pId,vals["time"][0]}
		fmt.Println(mess.Name+": "+mess.Text+"\n"+mess.Time+"\n")
	}
}

func main() {
	var id [32]byte
	mess = Message{"Welcome to Chat on the Go!","ServerAdmin",id,time.Now().String()}
	http.HandleFunc("/", speak)
	http.HandleFunc("/send", send)
	fmt.Println("Everything is working and ready...")
	http.ListenAndServe(":8080", nil)
}