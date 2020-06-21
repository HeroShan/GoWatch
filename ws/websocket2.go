package ws

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:1997", "http service address")

func Serve() {
	flag.Parse()
	hub := NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
