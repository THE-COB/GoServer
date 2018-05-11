package main

import(
	"fmt";
	"net/http";
	"net/url";
	"os";
	"os/exec";
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
		os.Create("./user.info")
		fmt.Println("It seems that you don't exist yet\nWhat is your name?")
		reader := bufio.NewReader(os.Stdin)
		name,_ := reader.ReadString('\n')
		ourBoi := Person{name, sha256.Sum256([]byte(time.Now().String()+name+string(rand.Intn(100))))}

		ioutil.WriteFile("./user.info", enc(ourBoi),'\n')
	}
	encBoi,err := ioutil.ReadFile("./user.info")
	err=err
	return decP(encBoi)
}

var mess Message
func checkDone(isDone *bool){
	reader := bufio.NewReader(os.Stdin)
	for true{
		status, _ := reader.ReadString('\n')
		if(status == "0\n"){
			*isDone = true
		} else if(status == "/t\n"){
			fmt.Println(mess.TimeSent)
		} else if(status == "/clear\n"){
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			cmd = exec.Command("cmd", "/c", "cls")
			cmd.Run()
		} else{
			p := getPerson()
			data := url.Values{}
			data.Add("sender", p.Name[:len(p.Name)-1])
			data.Add("sender", string(p.Id[:32]))
			data.Set("message", status[:len(status)-1])
			data.Set("time", time.Now().Format(time.RFC3339))
			resp, err := http.PostForm("http://localhost:8080/send", data)
			err=err
			defer resp.Body.Close()
		}
	}
}

func getMessagae() Message{
	resp, err := http.Get("http://localhost:8080")
	if(err != nil){
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var bMessage BasicMessage
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &bMessage)
	time,_ := time.Parse(time.RFC3339 ,bMessage.Time)
	p := Person{bMessage.Name, bMessage.Id}
	message := Message{p, bMessage.Text, time}
	return message
}

func printMess(d time.Duration){
	fmt.Println(getMessagae().Text)
	var lastMessage Message
	for true{
		time.Sleep(d)
		mess = getMessagae()
		if(lastMessage.TimeSent != mess.TimeSent && mess.Sender.Id != getPerson().Id){
			lastMessage = mess
			fmt.Println("\t"+mess.Sender.Name+": "+mess.Text)
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