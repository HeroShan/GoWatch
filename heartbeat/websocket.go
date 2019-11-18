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
	var sd,send_msg string
	for{
		n, err := ws.Read(msg)
		if err != nil ||n==0 {
			log.Fatal(err)
			break
		}
		sd += string(msg[:n])
		send_msg += "[" + string(msg[:n]) + "]"
	}
	cc,_ := base64.StdEncoding.DecodeString(sd)
	fmt.Printf("Service---Receive:----------- %v--%v\n", cc,time.Now())

	
	m, err := ws.Write([]byte(send_msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Service---Send:----------- %s--%v\n", msg[:m],time.Now())
}

func Server() {
	http.Handle("/echo", websocket.Handler(EchoHandler))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

var origin string
var url string
var i int

func Client(iplist []string) {
	fmt.Println(iplist)
	chanMax := len(iplist)
	writeChan := make(chan int, chanMax)
	for {
		i++
		time.Sleep(1 * time.Second)
		if i == 999999999 {
			i = 0
			i++
		}
		writeChan <- i
		go Send(writeChan, iplist)

	}

}

func Send(writeChan chan int, iplist []string) {
	for _, v := range iplist {
		origin = "http://" + strings.TrimSpace(v) + ":80/"
		url = "ws://" + strings.TrimSpace(v) + ":80/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err == nil {
			message := []byte(base64.StdEncoding.EncodeToString([]byte(v)))
			ws.Write(message)
			fmt.Printf("Client---Send: %s---%v--ip:--%v\n", message,time.Now(),v)
			<-writeChan
			var msg = make([]byte, 512)
			m, _ := ws.Read(msg)
			fmt.Printf("Client---Receive: %s---%v--ip:--%v\n", msg[:m],time.Now(),v)
			
		}
	}
}
