package main

import(
	"fmt";
	"net/http";
)

var text string
func speak(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, text)
}

func send(w http.ResponseWriter, r *http.Request){
	if(r.Method == "POST"){
		r.ParseForm()
		message := r.PostFormValue("message")
		text = message
	}
}

func main() {
	text = "Welcome to Chat on the Go!"
	http.HandleFunc("/", speak)
	http.HandleFunc("/send", send)
	fmt.Println("Everything is working and ready...")
	http.ListenAndServe(":8080", nil)
}