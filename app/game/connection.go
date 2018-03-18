package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	user models.User
	send chan []byte
}

func (client *Client) handleSendMessage() {
	for message := range client.send {
		client.conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (client *Client) handleHub(hub *Hub) {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			hub.unregister <- client
			break
		}

		var messageJSON map[string]interface{}
		json.Unmarshal(message, &messageJSON)
		if messageJSON["Type"] == "joinRoom" {
			type Name struct{ Name string }
			type Info struct {
				Payload int
			}
			var info Info
			json.Unmarshal(message, &info)

			hub.joinRoom <- JoinRoom{client, info.Payload}
			break
		}
	}
}

type JoinRoom struct {
	client *Client
	roomId int
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	joinRoom   chan JoinRoom
	rooms      map[int]*Room
}

func newHub() *Hub {
	hub := &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		joinRoom:   make(chan JoinRoom),
		clients:    make(map[*Client]bool),
		rooms:      make(map[int]*Room),
	}
	hub.rooms[0] = NewRoom(0, "rom0")
	hub.rooms[1] = NewRoom(1, "rom1")
	go hub.rooms[0].run(hub)
	go hub.rooms[1].run(hub)
	return hub
}

func (hub *Hub) run() {
	for {
		select {
		case newClient := <-hub.register:
			for client := range hub.clients {
				client.send <- NewUserMessage(newClient.user)
				newClient.send <- NewUserMessage(client.user)
			}
			fmt.Println("client joined hub ID=" + fmt.Sprint(newClient.user.ID))
			hub.clients[newClient] = true
			go newClient.handleHub(hub)
		case leavingClient := <-hub.unregister:
			delete(hub.clients, leavingClient)
			close(leavingClient.send)
			for client := range hub.clients {
				client.send <- DeleteUserMessage(leavingClient.user)
			}
			fmt.Println("client disconedted ID=" + fmt.Sprint(leavingClient.user.ID))
		case joinRoom := <-hub.joinRoom:
			delete(hub.clients, joinRoom.client)
			hub.rooms[joinRoom.roomId].Join <- joinRoom.client
			fmt.Println("Client ID=", joinRoom.client.user.ID, "Joined room ID=", joinRoom.roomId)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	userID := r.URL.Query()["user"]
	if userID == nil {
		conn.Close()
		return
	}
	var user models.User
	db.Db().Find(&user, userID)
	if user.ID == 0 {
		conn.Close()
		return
	}

	client := Client{conn, user, make(chan []byte, 5)}
	go client.handleSendMessage()
	hub.register <- &client
}

func RegisterRoutes(router *mux.Router) {
	hub := newHub()
	go hub.run()
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}
