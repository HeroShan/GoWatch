package ws

import (
	"GoWatch/createToken"
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		log.Printf("mesRead---:%v\n", string(message))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	qconn := MqConn()
	q, _ := qconn.QueueDeclare( //连接队列消费者模式
		"wsMessage", //Queue name
		true,        //durable
		false,
		false,
		false,
		nil,
	)
	for {
		select {
		case wstoken, ok := <-c.send:
			log.Printf("wstoken---:%v\n", string(wstoken))

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, NextWritererr := c.conn.NextWriter(websocket.TextMessage)
			if NextWritererr != nil {
				log.Printf("NextWritererr:%#v\n", NextWritererr)
			}
			expire := createToken.IsLogin(string(wstoken))
			log.Println(expire)
			if expire < 0 { //判断客户端连接是否正确
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			msgs, readMqMsgErr := qconn.Consume(
				q.Name,
				"",
				true, //Auto Ack
				false,
				false,
				false,
				nil,
			)
			if readMqMsgErr != nil {
				log.Printf("readMqMsgErr---:%#v\n", readMqMsgErr)
			}
			log.Printf("***********msgs---:%#v\n", msgs)
			w.Write(newline)
			w.Write([]byte("123"))
			n := len(c.send)
			log.Printf("***********n---:%#v\n", n)
			for msg := range msgs {
				for i := 0; i < n; i++ {
					log.Printf("Body---:%#v\n", string(msg.Body))
					w.Write(newline)
					w.Write(msg.Body)
					<-c.send
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
