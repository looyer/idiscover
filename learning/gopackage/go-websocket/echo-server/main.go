package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	log.Printf("~~~~~~~~~ echo-server ~~~~~~~~~")

	wsport := int(8080)
	flag.IntVar(&wsport, "wsport", wsport, "websocket listen port")
	flag.Parse()

	http.HandleFunc("/echo/wsEcho", wsEcho)
	http.ListenAndServe(fmt.Sprintf(":%v", wsport), nil)
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(req *http.Request) bool {
			return true
		},
	}
)

func wsEcho(w http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			log.Printf("ws-upgrade err:%v remote:%v", err, r.RemoteAddr)
		}
		go echo(conn)

	} else {
		fmt.Fprintf(w, "please use websocket!")
	}
}

func echo(conn *websocket.Conn) {
	defer conn.Close()

	cliaddr := conn.RemoteAddr().String()

	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 60))
		t, buffer, err := conn.ReadMessage()
		if os.IsTimeout(err) {
			log.Printf("echo i/o timeout err:%v cliaddr:%v", err, cliaddr)
			return
		}
		if err != nil {
			log.Printf("echo read err:%v cliaddr:%v", err, cliaddr)
			return
		}

		if t == websocket.TextMessage {
			log.Printf("echo recv cliaddr:%v text:%v", cliaddr, string(buffer))
		}
		conn.WriteMessage(t, buffer)
	}
}
