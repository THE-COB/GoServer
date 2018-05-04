package main

import(
	"fmt";
	"net/http";
	"net/url";
	"os";
	"os/exec"
	"bufio";
	"io/ioutil";
	"time";
	"crypto/sha256";
	"math/rand";
	"encoding/json"
)

func getPerson() Person{
	_,noFile := os.Stat("./user.info")
	if noFile != nil{
		newFile,err := os.Create("./user.info")
		newFile=newFile
		fmt.Println(err)
		fmt.Println("It seems that you don't exist yet\nWhat is your name?")
		reader := bufio.NewReader(os.Stdin)
		name,_ := reader.ReadString('\n')
		ourBoi := Person{name, sha256.Sum256([]byte(name+string(rand.Intn(100))))}

		err = ioutil.WriteFile("./user.info", enc(ourBoi),'\n')
	}
	encBoi,err := ioutil.ReadFile("./user.info")
	err=err
	return decP(encBoi)
}

func checkDone(isDone *bool){
	reader := bufio.NewReader(os.Stdin)
	for true{
		status, _ := reader.ReadString('\n')
		if(status == "0\n"){
			*isDone = true
		} else if(status == "clear\n"){
			exec.Command("clear")
		} else{
			p := getPerson()
			data := url.Values{}
			data.Add("sender", p.Name)
			data.Add("sender", string(p.Id[:32]))
			data.Set("message", status[:len(status)-1])
			data.Set("time", time.Now().String())
			resp, err := http.PostForm("http://localhost:8080/send", data)
			err=err
			defer resp.Body.Close()
		}
	}
}

func getMessagae() string{
	resp, err := http.Get("http://localhost:8080")
	if(err != nil){
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var message string
	var mess []byte
	json.Unmarshal(mess, resp.Body)
	fmt.Println(mess)
	return message
}

func printMess(d time.Duration){
	var lastMessage string
	for true{
		time.Sleep(d)
		mess := getMessagae()
		if(true){
			lastMessage = mess
			lastMessage=lastMessage
			fmt.Println(mess)
		}

	}
}

func main(){
	isDone := false
	go printMess(1*time.Second)
	go checkDone(&isDone)
	for true{
		if(isDone){
			os.Exit(3)
		}
	}
}