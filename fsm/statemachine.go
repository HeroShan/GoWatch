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
	conn, _ := http.Get("http://wanter.work:8080/fmsgetip")
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
	ip := strings.Split(filestr,"=")
	Hplist := strings.Split(ip[1],",")
	for k,v := range Hplist{
		v = strings.TrimSpace(v)
		if Hplist[k] == locIp{
			go heartbeat.Server()
		}
	}
	for k,v := range Hplist{
		v = strings.TrimSpace(v)
		if v == locIp {
			Hplist = append(Hplist[:k],Hplist[k+1:]...)
		}
	}
	heartbeat.Client(Hplist,locIp)
}

func main(){
	HeartBeat()
	//heartbeat.Auth()
}
