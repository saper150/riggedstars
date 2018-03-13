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
	Leave   chan *Client
	Join    chan *Client
}

func (room *Room) run(hub *Hub) {
	for {
		select {
		case message := <-room.Message:
			for client := range room.Clients {
				if client.user != message.From.user {
					client.conn.WriteMessage(websocket.TextMessage, CreateTextMessage(message.Message))
				}
			}
		case client := <-room.Leave:
			delete(room.Clients, client)
			hub.register <- client
		case client := <-room.Join:
			room.Clients[client] = true
			go client.handleRoom(room)
		}
	}
}

func (client *Client) handleRoom(room *Room) {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading room")
			room.Leave <- client
			break
		}

		var messageJSON map[string]interface{}
		json.Unmarshal(message, &messageJSON)

		if messageJSON["Type"] == "text" {
			fmt.Println("text message from ID=", client.user.ID)
			type Name struct{ Name string }
			type Info struct {
				Payload string
			}
			var info Info
			json.Unmarshal(message, &info)
			room.Message <- ClientMessage{client, info.Payload}
		} else if messageJSON["Type"] == "leaveRoom" {
			room.Leave <- client
			break
		}

	}
}
