package heartbeat

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"encoding/base64" 
	"golang.org/x/net/websocket"
)

//wingman WINGMAN  ↑↑ ↓↓  ←→  ←→ BABA卒
func EchoHandler(ws *websocket.Conn) {	
    msg := make([]byte, 512)	
    n, err := ws.Read(msg)	
    if err != nil {	
        log.Fatal(err)	
	}	
	clientip,_ := base64.StdEncoding.DecodeString(string(msg[:n]))
	fmt.Println(string(clientip))
	
	
	Hplist := GetConfIp()
	for _,v := range Hplist{
		v = strings.TrimSpace(v)
		if v == string(clientip){
			send_msg := "wingman WINGMAN  ↑↑ ↓↓  ←→  ←→ BABA卒"+v	
			_, err := ws.Write([]byte(send_msg))	
		    if err != nil {	
		        log.Fatal(err)	
		    }	
		    //fmt.Printf("Send:******* %s\n", msg[:m])
		}	
	}
	
}


func Server() {
	http.Handle("/echo", websocket.Handler(EchoHandler))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func Client(iplist []string,locIp string) {
	chanMax := len(iplist)
	writeChan := make(chan string, chanMax)
	for {
		for _,fsmIp := range iplist{
			writeChan<-fsmIp
			go Send(writeChan,locIp)
		}
		
		time.Sleep(3 * time.Second)
	}

}

func Send(writeChan chan string,locIp string) {
	var origin string
	var url string
	v := <-writeChan
		origin = "http://" + strings.TrimSpace(v) + ":80/"
		url = "ws://" + strings.TrimSpace(v) + ":80/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err != nil{
			fmt.Println(locIp)
		}else{
			var msg = make([]byte, 512)
			//message := []byte(base64.StdEncoding.EncodeToString([]byte(locIp)))
			msint,err1:=ws.Read(msg)
			if err1 != nil{
				fmt.Println(err1,"---------------------------",msint,"-----------------",locIp)
			}
			
		}
}

func CheckBeat(receivedata string,now time.Time,ip string){
	fmt.Printf("%s,%s,%s\n",receivedata,now,ip)
}