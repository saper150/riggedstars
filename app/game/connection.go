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
	//TODO: rooms auto creation
	hub.rooms[0] = newRoom(0, "rom0", 4)
	hub.rooms[1] = newRoom(1, "rom1", 4)
	go hub.rooms[0].run(hub)
	go hub.rooms[1].run(hub)
	return hub
}

func newRoom(id int, name string, maxClients int) *Room {
	return &Room{
		ID:         id,
		Clients:    make(map[*Client]bool),
		Name:       name,
		Message:    make(chan ClientMessage),
		Leave:      make(chan *Client),
		Join:       make(chan *Client),
		MaxClients: maxClients,
	}
}

func (hub *Hub) run() {
	for {
		select {
		case newClient := <-hub.register:
			for client := range hub.clients {
				client.conn.WriteMessage(websocket.TextMessage, NewUserMessage(newClient.user))
				newClient.conn.WriteMessage(websocket.TextMessage, NewUserMessage(client.user))
			}
			fmt.Println("client joined hub ID=" + fmt.Sprint(newClient.user.ID))
			hub.clients[newClient] = true
			go newClient.handleHub(hub)
		case leavingClient := <-hub.unregister:
			delete(hub.clients, leavingClient)
			for client := range hub.clients {
				client.conn.WriteMessage(websocket.TextMessage, DeleteUserMessage(leavingClient.user))
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

	client := Client{conn, user}
	hub.register <- &client
}

func RegisterRoutes(router *mux.Router) {
	hub := newHub()
	go hub.run()
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	router.HandleFunc("/roomList", hub.getRoomList).Methods("GET")
}

type roomInfo struct {
	ID           int
	ClientsCount int
	Name         string
	MaxClients   int
}

func (hub *Hub) getRoomList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	rooms := make([]roomInfo, len(hub.rooms))

	for i := 0; i < len(hub.rooms); i++ {
		rooms[i] = roomInfo{ID: hub.rooms[i].ID, Name: hub.rooms[i].Name, ClientsCount: len(hub.rooms[i].Clients), MaxClients: hub.rooms[i].MaxClients}
	}
	js, err := json.Marshal(rooms)
	if err == nil {
		w.Write([]byte(js))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
