package ws
// 1.开启端口监听
// 2.开启ws服务
// 3.接受队列数据再发送ws推送


import (  
	"log" 
	"github.com/gorilla/websocket" 
	"net/http"
 )  
 
 
// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//即时消息传递
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// ws建立服务
	wsConn, _ := wsUpgrader.Upgrade(w, r, nil)
	//rabbitmq读取消息队列
	conn := MqConn()
	q, MqConnErr := conn.QueueDeclare(	//连接队列消费者模式
        "wsMessage", //Queue name
        true, //durable
        false,
        false,
        false,
        nil,
	)
	conn.Consume(		//获取队列消息
        q.Name,
        "",
        true,  //Auto Ack
        false,
        false,
        false,
        nil,
    )
	if MqConnErr != nil{
		log.Printf("Mq connect failed:%v\n",MqConnErr)
	}
	defer wsConn.Close()
	for{
		mt, message, err := wsConn.ReadMessage()
		if err != nil{
			break 
		}
		if message != nil{
			err = wsConn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
			}
			log.Printf("clientMsg:%v\n",string(message))
		}
		
	}
	
	
	// for msg := range msgs{
	// 	if len(msg.Body)>0{
	// 				//把消息广播到正在ws客户端连接
	// 				err = wsConn.WriteMessage(websocket.TextMessage, msg.Body)
	// 				if err != nil {
	// 					log.Println("write:", err)
	// 				}
	// 		}
	// }
		
	
}
 
 func Serve() { 
	http.HandleFunc("/ws", wsHandler)
	log.Print(http.ListenAndServe("127.0.0.1:1997", nil))
 }