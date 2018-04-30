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
	"encoding/gob";
	"bytes";
	"crypto/sha256";
	"math/rand"
)

func getPerson() Person{
	_,noFile := os.Stat("./user.info")
	fmt.Println(noFile)
	if noFile != nil{
		newFile,err := os.Create("./user.info")
		newFile=newFile
		fmt.Println(err)
		fmt.Println("It seems that you don't exist yet\nWhat is your name?")
		reader := bufio.NewReader(os.Stdin)
		name,_ := reader.ReadString('\n')
		ourBoi := Person{name, sha256.Sum256([]byte(name+string(rand.Intn(100))))}

		err = ioutil.WriteFile("./user.info", encPerson(ourBoi),'\n')
	}
	encBoi,err := ioutil.ReadFile("./user.info")
	err=err
	return decPerson(encBoi)
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
			data := url.Values{}
			data.Set("message", status[:len(status)-1])
			enc,_ := ioutil.ReadFile("./user.info")
			data.Add("person", string(enc))
			resp, err := http.PostForm("http://localhost:8080/send", data)
			err=err
			defer resp.Body.Close()
		}
	}
}

func encPerson(boi Person) []byte{
	var net bytes.Buffer
	enc := gob.NewEncoder(&net)
	err := enc.Encode(boi)
	err = err
	return net.Bytes()
}
func decPerson(enc []byte) Person{
	var boi Person
	dec := gob.NewDecoder(bytes.NewReader(enc))
	err := dec.Decode(&boi)
	err=err
	return boi
}

func getMessagae() string{
	resp, err := http.Get("http://localhost:8080")
	if(err != nil){
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var message string
	for i := range(body){
		message+=string(body[i])
	}
	return message
}

func printMess(d time.Duration){
	lastMessage := ""
	for true{
		time.Sleep(d)
		txt := getMessagae()
		if(txt != lastMessage){
			lastMessage = txt
			fmt.Println(txt)
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