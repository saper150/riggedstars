package game

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type ClientMessage struct {
	From    *Client
	Message string
}

type Room struct {
	ID      int
	Name    string
	Clients map[*Client]bool
	Message chan ClientMessage
}

func (room *Room) run() {
	for {
		message := <-room.Message
		for client := range room.Clients {
			if client.user != message.From.user {
				client.conn.WriteMessage(websocket.TextMessage, CreateTextMessage(message.Message))
			}
		}
	}
}

func (client *Client) handleRoom(room *Room) {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("left")
		}

		var messageJSON map[string]interface{}
		json.Unmarshal(message, &messageJSON)

		if messageJSON["Type"] == "text" {
			type Name struct{ Name string }
			type Info struct {
				Payload string
			}
			var info Info
			json.Unmarshal(message, &info)
			room.Message <- ClientMessage{client, info.Payload}
		}
	}
}
