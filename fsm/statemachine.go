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
	tmpMap := make(map[string]bool,len(Hplist))
	Siplist := make([]string,0)
	var allowSer bool
	for _,v := range Hplist{
		if v == locIp{
			allowSer = true
		}
		v = strings.TrimSpace(v)
		if tmpMap[v] == false && v !=locIp{
		   tmpMap[v] =	true
		   Siplist = append(Siplist,v)
		}
	}
	if allowSer == true{
		go heartbeat.Server()
	}
	//fmt.Println(allowSer,Siplist,locIp)
	heartbeat.Client(Siplist,locIp)
}

func main(){
	HeartBeat()
}
