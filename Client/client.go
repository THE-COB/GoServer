package main

import(
	"fmt";
	"net/http";
	"os";
	"bufio";
	"io/ioutil"
)

func checkDone(isDone *bool){
	reader := bufio.NewReader(os.Stdin)
	for true{
		status, _ := reader.ReadString('\n')
		if(status == "0\n"){
			*isDone = true
		}
	}
}

func main(){
	isDone := false
	resp, err := http.Get("http://localhost:8080")
	if(err != nil){
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	go checkDone(&isDone)
	for true{
		if(isDone){
			os.Exit(3)
		}
	}
}