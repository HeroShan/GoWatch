package heartbeat

<<<<<<< HEAD
import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

func EchoHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive:------------------- %s%v\n", msg[:n], time.Now())

	send_msg := "[" + string(msg[:n]) + "]"
	m, err := ws.Write([]byte(send_msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send:------------------- %s%v\n", msg[:m], time.Now())
=======
import(
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "time"
    "strings"
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
>>>>>>> 9f20dff7adcf6f03709f94a3017f759a6feceb2a
}

func Server() {
	http.Handle("/echo", websocket.Handler(EchoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

var origin string
var url string
var i int

func Client(iplist []string) {
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

<<<<<<< HEAD
func Send(writeChan chan int, iplist []string) {
	for _, v := range iplist {
		origin = "http://" + strings.TrimSpace(v) + ":8080/"
		url = "ws://" + strings.TrimSpace(v) + ":8080/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err == nil {
			message := []byte("2019")

			ws.Write(message)
			fmt.Printf("Send: %s\n", message)
			<-writeChan
			var msg = make([]byte, 512)
			m, _ := ws.Read(msg)
			fmt.Printf("Receive: %s\n", msg[:m])
		}
	}
=======
func Send(writeChan chan int,iplist []string){
        for _,v := range iplist{
            origin = "http://"+strings.TrimSpace(v)+":8080/"
            url = "ws://"+strings.TrimSpace(v)+":8080/echo"
            ws, err := websocket.Dial(url, "", origin)
            if err == nil {
                message := []byte("2019")
                
                ws.Write(message)
                fmt.Printf("Send: %s\n", message)
                <-writeChan
                var msg = make([]byte, 512)
                m, _ := ws.Read(msg)
                fmt.Printf("Receive: %s\n", msg[:m])
            }
        }
>>>>>>> 9f20dff7adcf6f03709f94a3017f759a6feceb2a
}
