package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade the HTTP connection to a WebSocket connection
var upgrader = websocket.Upgrader{
	// CheckOrigin allows all connections by returning true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler function for WebSocket connections
func handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Infinite loop to read and write messages
	for {
		// Read message from WebSocket
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Write message back to WebSocket
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	// Set the handler for the "/ws" endpoint
	http.HandleFunc("/ws", handler)
	fmt.Println("Server started at :8080")

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error:", err)
	}
}
