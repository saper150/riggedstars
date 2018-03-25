package game

import (
	"encoding/json"
	"log"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type genericCommand struct {
	Type    string
	Payload interface{}
}

type Client struct {
	conn        *websocket.Conn
	user        models.User
	sendMessage chan interface{}
	listen      chan genericCommand
	disconect   chan interface{}
}

func (client *Client) serveMessages() {
	for message := range client.sendMessage {
		client.conn.WriteJSON(message)
	}
}

func (client *Client) serveListen() {

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			close(client.sendMessage)
			close(client.listen)
			client.disconect <- ""
			break
		}
		var genericCommand genericCommand
		json.Unmarshal(message, &genericCommand)
		client.listen <- genericCommand

	}
}

type Hub struct {
	rooms map[int]*Room
}

func newHub() *Hub {
	hub := &Hub{
		rooms: make(map[int]*Room),
	}
	//TODO: rooms auto creation
	hub.rooms[0] = newRoom(0, "room0", 4)
	hub.rooms[1] = newRoom(1, "room1", 4)
	go hub.rooms[0].run(hub)
	go hub.rooms[1].run(hub)
	return hub
}

func newRoom(id int, name string, maxClients int) *Room {
	return &Room{
		ID:       id,
		Clients:  make(map[*Client]bool),
		Name:     name,
		Commands: make(chan clientCommand),
		Leave:    make(chan *Client),
		Join:     make(chan *Client),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func joinRoom(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	if r.URL.Query()["user"] == nil || r.URL.Query()["roomId"] == nil {
		conn.Close()
		return
	}

	var user models.User
	db.Db().Find(&user, r.URL.Query()["user"])
	if user.ID == 0 {
		conn.Close()
		return
	}

	client := Client{
		conn:        conn,
		user:        user,
		sendMessage: make(chan interface{}, 5),
		listen:      make(chan genericCommand),
		disconect:   make(chan interface{}),
	}
	go client.serveListen()
	go client.serveMessages()

	var roomID int
	roomID, err = strconv.Atoi(r.URL.Query().Get("roomId"))

	if err != nil {
		conn.Close()
		return
	}

	room, ok := hub.rooms[roomID]
	if !ok {
		conn.Close()
		return
	}
	room.Join <- &client
}

func RegisterRoutes(router *mux.Router) {
	hub := newHub()
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		joinRoom(hub, w, r)
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
