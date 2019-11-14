package heartbeat

import(
	"golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "time"
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

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/echo"
 
func Client(iplist []string) {
    for {
        time.Sleep(2*time.Second)
        ws, err := websocket.Dial(url, "", origin)
        if err != nil {
            log.Fatal(err)
        }
        message := []byte("asdasdasd")
        _, err = ws.Write(message)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Send: %s\n", message)
    
        var msg = make([]byte, 512)
        m, err := ws.Read(msg)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Receive: %s\n", msg[:m])
    
        defer ws.Close()//关闭连接
    }
}