package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akwrs/lunch-delivery/types"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Receiver working fine")

	receiver := &DataReceiver{
		msgCh: make(chan types.Rider, 128),
	}

	http.HandleFunc("/ws", receiver.handleWS)
	log.Fatal(http.ListenAndServe(":30000", nil))
}

type DataReceiver struct {
	msgCh chan types.Rider
	conn  *websocket.Conn
}

// WebSocket handler
func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming WebSocket upgrade request")

	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	fmt.Println("New WebSocket connection accepted from:", r.RemoteAddr)
	dr.conn = conn

	go dr.wsReceiveLoop()
}

// Receive and print JSON messages
func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("Waiting for Rider messages...")

	for {
		var data types.Rider
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			return // close the loop on error (e.g., client disconnect)
		}

		fmt.Println("Received Rider:", data)

		// Optionally push to channel
		select {
		case dr.msgCh <- data:
		default:
			// channel full or unused
		}
	}
}
