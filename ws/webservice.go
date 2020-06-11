package ws
// 1.开启端口监听
// 2.开启ws服务
// 3.接受队列数据再发送ws推送


import (  
	"log" 
	"github.com/gorilla/websocket" 
	"net/http"
	_"time"
	"GoWatch/createToken"
	)  
 
 
// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Mqmsg  chan []byte

//即时消息传递
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// ws建立服务
	wsConn, _ := wsUpgrader.Upgrade(w, r, nil)
	conn := MqConn()
			q, _ := conn.QueueDeclare(	//连接队列消费者模式
				"wsMessage", //Queue name
				true, //durable
				false,
				false,
				false,
				nil,
			)
	
		defer wsConn.Close()
	for{
		mt, wstoken, err := wsConn.ReadMessage()
		expire := createToken.IsLogin(string(wstoken))
		if err != nil{
			break 
		}
		if expire > 0{
			
			msgs, _ := conn.Consume(
				q.Name,
				"",
				true,  //Auto Ack
				false,
				false,
				false,
				nil,
			)
			for msg := range msgs{
				if len(msg.Body)>0{
					wsConn.WriteMessage(mt, msg.Body)
					}
			}
			
		}
		
	}		
	
}

 
func Serve() { 
	http.HandleFunc("/ws", wsHandler)
	log.Print(http.ListenAndServe("127.0.0.1:1997", nil))
 }