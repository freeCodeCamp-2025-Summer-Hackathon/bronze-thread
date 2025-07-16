package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for development purposes.
		// You might want to implement more restrictive checks in production.
		return true
	},
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer ws.Close()

	log.Println("Client connected")

	for {
		// Read message from browser
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// Print the message to the console
		log.Printf("Received: %s", msg)

		// Write message back to browser
		if err := ws.WriteMessage(msgType, msg); err != nil {
			log.Println("Failed to write message:", err)
			break
		}
	}
}
