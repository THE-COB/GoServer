package main

import(
	"fmt";
	"net/http"
)

func sayHi(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hi</h1>")
}

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/", sayHi)
	http.ListenAndServe(":8080", nil)
}