package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/websocket"
	"github.com/gin-gonic/gin"
)

// ServeWs handles websocket requests from the peer.
// It requires a reference to the central Hub.
func ServeWs(hub *websocket.Hub, c *gin.Context) {
	// Get user ID from the JWT token (set by the IsAuthorized middleware)
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get room ID from the URL path parameter
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Room ID"})
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new client instance for this connection
	client := &websocket.Client{
		Hub:    hub,
		Conn:   conn,
		Send:   make(chan models.ChatMessage, 256), // Buffered channel
		UserID: userID.(uint),
		RoomID: uint(roomID),
	}

	// Register the new client with the hub
	client.Hub.Register <- client

	// Start the concurrent read and write pumps
	go client.WritePump()
	go client.ReadPump()
}
