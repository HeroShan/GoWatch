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
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//即时消息传递
func wsHandler(resp http.ResponseWriter, req *http.Request) {
	// ws建立服务
	c, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	//rabbitmq读取消息队列
	conn := Conn()
	q, MqConnErr := conn.QueueDeclare(	//连接队列消费者模式
        "wsMessage", //Queue name
        true, //durable
        false,
        false,
        false,
        nil,
	)
	msgs, _ := conn.Consume(		//获取队列消息
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
	
	defer c.Close()
	for msg := range msgs{
		if len(msg.Body)>0{
					//把消息广播到正在ws客户端连接
					err = c.WriteMessage(websocket.TextMessage, msg.Body)
					if err != nil {
						log.Println("write:", err)
					}
					
			}
	}
		
	
}
 
 func Serve() { 
	http.HandleFunc("/ws", wsHandler)
	log.Print(http.ListenAndServe(":1997", nil))
 }