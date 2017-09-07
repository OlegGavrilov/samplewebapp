package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // channel to store messages

// Configure the upgrader
var wsupgrader = websocket.Upgrader{}

// Message - struct to hold message details
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	// convert http connection to websocket
	connection, err := wsupgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to convert connection to websocket: %+v", err)
		return
	}

	// save client
	clients[connection] = true

	for {
		var message Message
		// deserialize json
		err := connection.ReadJSON(&message)
		if err != nil {
			log.Printf("[%v] Error in parsing Message JSON: %v", connection.RemoteAddr().String(), err)
			break
		}
		// send message to channel
		broadcast <- message
	}

	defer connection.Close()
}

func handleMessages() {
	for {
		// take next message from channel
		message := <-broadcast

		// broadcast to all connections in clients map
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
