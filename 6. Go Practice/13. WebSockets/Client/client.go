package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the WebSocket server URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Printf("Connecting to %s\n", u.String())

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send a message to the WebSocket server
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, Server!"))
	if err != nil {
		log.Fatal("Write error:", err)
	}

	// Read a message from the WebSocket server
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Printf("Received: %s\n", message)
}
