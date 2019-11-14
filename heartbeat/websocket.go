package heartbeat

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
    fmt.Printf("Receive:------------------- %s\n", msg[:n])
 
    send_msg := "[" + string(msg[:n]) + "]"
    m, err := ws.Write([]byte(send_msg))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Send:------------------- %s\n", msg[:m])
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
var url  string
 
func Client(iplist []string) {

    for {
        for _,v := range iplist{
            origin = "http://"+strings.TrimSpace(v)+":8080/"
            url = "ws://"+strings.TrimSpace(v)+":8080/echo"
            ws, _ := websocket.Dial(url, "", origin)
            message := []byte("asdasdasd")
            time.Sleep(1*time.Second)
            ws.Write(message)
            fmt.Printf("Send: %s\n", message)
        
            var msg = make([]byte, 512)
            m, _ := ws.Read(msg)
            fmt.Printf("Receive: %s\n", msg[:m])
        
        }
        
    }
}