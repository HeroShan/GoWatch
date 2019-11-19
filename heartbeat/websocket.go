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

func EchoHandler(ws *websocket.Conn) {	
    msg := make([]byte, 512)	
    n, err := ws.Read(msg)	
    if err != nil {	
        log.Fatal(err)	
    }	
	fmt.Printf("Receive:---------------- %s\n", msg[:n])
	send_msg := "[" + string(msg[:n]) + "]"	
    m, err := ws.Write([]byte(send_msg))	
    if err != nil {	
        log.Fatal(err)	
    }	
    fmt.Printf("Send:------------------- %s\n", msg[:m])	
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
		
		time.Sleep(1 * time.Second)
	}

}

func Send(writeChan chan string,locIp string) {
	var origin string
	var url string
		v := <-writeChan
		origin = "http://" + strings.TrimSpace(v) + ":80/"
		url = "ws://" + strings.TrimSpace(v) + ":80/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err == nil {
			message := []byte(base64.StdEncoding.EncodeToString([]byte(locIp)))
			ws.Write(message)
			fmt.Printf("Client---Send: %s---%v--ip:--%v\n", message,time.Now(),v)
			var msg = make([]byte, 512)
			m, _ := ws.Read(msg)
			fmt.Printf("Client---Receive: %s---%v--ip:--%v\n", msg[:m],time.Now(),v)
		}
}