package heartbeat

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
	"encoding/base64" 
	"golang.org/x/net/websocket"
)
//wingman WINGMAN  ↑↑ ↓↓  ←→  ←→ BABA
func EchoHandler(ws *websocket.Conn) {	
    msg := make([]byte, 512)	
    n, err := ws.Read(msg)	
    if err != nil {	
        log.Fatal(err)	
	}	
	clientip,_ := base64.StdEncoding.DecodeString(string(msg[:n]))
	file,_ := os.Open("../fsm/fsm.config")
	data := make([]byte,512)
	var filestr string
	for{
		count, _ := file.Read(data)
		filestr += string(data[:count])
		if count == 0{
			break
		}
	}
	ip := strings.Split(filestr,"=")
	Hplist := strings.Split(ip[1],",")
	for _,v := range Hplist{
		v = strings.TrimSpace(v)
		if v == string(clientip){
			
			send_msg := base64.StdEncoding.EncodeToString([]byte("wingman WINGMAN  ↑↑ ↓↓  ←→  ←→ BABA卒"+string(clientip)))	
			_, err := ws.Write([]byte(send_msg))	
		    if err != nil {	
		        log.Fatal(err)	
		    }	
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
var failedIp = make(chan string,1)
func Client(iplist []string,locIp string) {
	chanMax := len(iplist)
	writeChan := make(chan string, chanMax)
	for {
		for _,fsmIp := range iplist{
			select{
			case <-failedIp :
				time.Sleep(1 * time.Second)
			default :
			writeChan<-fsmIp
			}
			go Send(writeChan,locIp)
			
		}
		  time.Sleep(1 * time.Second)
	}

}

func Send(writeChan chan string,locIp string) {
	var origin string
	var url string
		sendip := <-writeChan
		origin = "http://" + strings.TrimSpace(sendip) + ":80/"
		url = "ws://" + strings.TrimSpace(sendip) + ":80/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err == nil {
			var msg = make([]byte, 512)
			fmt.Println("asddddddddddddqqqqqqqqqqqqqqqqqq")
			m, err1 := ws.Read(msg); if err1 !=nil{
				fmt.Println("ccccccccccccccccccccc",err1)
			}else{
				fmt.Println("ccccccccccccccccccccasdasdasdadc",m)
			}
			fmt.Println("asdddddddddddd",m)
			if m == 0{
				fmt.Println("11111")
			message := []byte(base64.StdEncoding.EncodeToString([]byte(locIp)))
			ws.Write(message)
			}
			clientip,_ := base64.StdEncoding.DecodeString(string(msg[:m]))
			if len(clientip)<54{
				failedIp <- sendip
			}
			ipDe := strings.Split(string(clientip),"卒")
			message := []byte(base64.StdEncoding.EncodeToString([]byte(ipDe[1])))
			ws.Write(message)
			CheckBeat(string(msg[:m]),time.Now(),sendip)
		}else{
			failedIp <- sendip
			fmt.Println(err)
		
		}
}

func CheckBeat(receivedata string,now time.Time,ip string){
	fmt.Printf("%s,%s,%s\n",receivedata,now,ip)
}