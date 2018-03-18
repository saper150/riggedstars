package game

import (
	"encoding/json"
	"fmt"
)

type RoomMessage struct {
	From    *RoomClient
	Message string
}

type RoomState int

const (
	wating      = iota
	playing     = iota
	disconected = iota
)

type RoomClient struct {
	State       RoomState
	Client      *Client
	GameMessage chan []byte
}

const (
	idle = iota
	game = iota
)

type Room struct {
	ID        int
	State     int
	Name      string
	Clients   map[*RoomClient]bool
	StartGame chan bool
	Message   chan RoomMessage
	Leave     chan *RoomClient
	Join      chan *Client
	Brodcast  chan []byte
}

func NewRoom(id int, name string) *Room {
	return &Room{
		ID:        id,
		Clients:   make(map[*RoomClient]bool),
		Name:      name,
		Message:   make(chan RoomMessage),
		Leave:     make(chan *RoomClient),
		Join:      make(chan *Client),
		State:     idle,
		StartGame: make(chan bool),
		Brodcast:  make(chan []byte),
	}
}

func (room *Room) run(hub *Hub) {
	var game *Game
	game = nil
	for {
		select {
		case message := <-room.Message:
			for roomClient := range room.Clients {
				if roomClient.Client.conn != message.From.Client.conn {
					roomClient.Client.send <- CreateTextMessage(message.Message)
				}
			}
		case <-room.StartGame:
			if game == nil {
				game = NewGame(room)
			}
		case roomClient := <-room.Leave:
			delete(room.Clients, roomClient)
			hub.register <- roomClient.Client
		case client := <-room.Join:
			roomClient := &RoomClient{wating, client, make(chan []byte)}
			room.Clients[roomClient] = true
			go roomClient.handleRoom(room)
		case message := <-room.Brodcast:
			for roomClient := range room.Clients {
				roomClient.Client.send <- message
			}
		}
	}
}

func (roomClient *RoomClient) handleRoom(room *Room) {
	for {
		_, message, err := roomClient.Client.conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading room")
			room.Leave <- roomClient
			break
		}

		var messageJSON map[string]interface{}
		json.Unmarshal(message, &messageJSON)

		if messageJSON["Type"] == "text" {
			fmt.Println("text message from ID=", roomClient.Client.user.ID)
			type Name struct{ Name string }
			type Info struct {
				Payload string
			}
			var info Info
			json.Unmarshal(message, &info)
			room.Message <- RoomMessage{roomClient, info.Payload}
		} else if messageJSON["Type"] == "startGame" {
			room.StartGame <- true
			break
		} else if messageJSON["Type"] == "leaveRoom" {
			room.Leave <- roomClient
			break
		}
	}
}
