package main

import(
	"net/http"
	"fmt"
	"os"
	"strings" 
	"GoWatch/heartbeat"
	"io/ioutil"


)


func Getlocalip() (string) {  
	conn, _ := http.Get("http://127.0.0.1:8080/fmsgetip")
    body, _ := ioutil.ReadAll(conn.Body)
	defer conn.Body.Close()
	return strings.Split(string(body), ":")[0]
}

func VoteTime(){

}

func HeartBeat(){
	file,err := os.Open("./fsm.config"); if err != nil{
		fmt.Println("file open fail",err)
	}
	data := make([]byte,512)
	var filestr string
	for{
		count, _ := file.Read(data)
		filestr += string(data[:count])
		if count == 0{
			break
		}
	}
	locIp := Getlocalip()
	fmt.Println(locIp)
	ip := strings.Split(filestr,"=")
	iplist := strings.Split(ip[1],",")
	for k,v := range iplist{
		v = strings.TrimSpace(v)
		if v == locIp {
			iplist = append(iplist[:k],iplist[k+1:]...)
		}
	}
	fmt.Println(iplist)
	go heartbeat.Server()
	heartbeat.Client(iplist)
}


func main(){
	HeartBeat()
	
}
