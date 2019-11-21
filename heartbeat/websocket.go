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
	clientip,eer := base64.StdEncoding.DecodeString(string(msg[:n]))
	if eer != nil {
		fmt.Println(eer)
		}
	//fmt.Printf("Server:******* %s\n", clientip)
	file,err := os.Open("../fsm/fsm.config"); if err != nil{
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
	ip := strings.Split(filestr,"=")
	Hplist := strings.Split(ip[1],",")
	for _,v := range Hplist{
		v = strings.TrimSpace(v)
		if v == string(clientip){
			send_msg := "wingman WINGMAN  ↑↑ ↓↓  ←→  ←→ BABA"	
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
		fmt.Println(v,err)
		if err == nil {
			message := []byte(base64.StdEncoding.EncodeToString([]byte(locIp)))
			ws.Write(message)
			//fmt.Printf("Client---Send: %s---%v--ip:--%v\n", message,time.Now(),v)
			var msg = make([]byte, 512)
			m, _ := ws.Read(msg)
			CheckBeat(string(msg[:m]),time.Now(),v)
			//fmt.Printf("Client---Receive: %T---%v--ip:--%v\n", msg[:m],time.Now(),v)
		}else{
			CheckBeat("failed",time.Now(),v)
		}
}

func CheckBeat(receivedata string,now time.Time,ip string){
	fmt.Printf("%s,%s,%s\n",receivedata,now,ip)
}