package ws
// 1.开启端口监听
// 2.开启ws服务
// 3.接受请求数据再发送ws推送


import (  
	"github.com/gin-gonic/gin"  
	"github.com/gorilla/websocket" 
	"net/http"
 )  
 
 
 var upGrader = websocket.Upgrader{  
	CheckOrigin: func (r *http.Request) bool {  
	   return true  
	},  
 }  
 
 //webSocket请求ping 返回pong  
 func Ping(c *gin.Context) {  
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)  
	if err != nil {  
	   return  
	}  
	var mes []byte
	defer ws.Close()  
	for {
		mes = []byte("123")
	   if string(mes) != ""{
		err = ws.WriteMessage(websocket.PingMessage,mes)  
		if err != nil {  
		   break  
		}
	   } 
	   //写入ws数据
	}  
 }  
 
 func Serve() { 
	//bindAddress := ":8081"  
	r := gin.Default() 
	 
	r.GET("/ping", Ping) 
 }