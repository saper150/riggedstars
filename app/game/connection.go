package game

import (
	"fmt"
	"log"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
}

const (
	waiting        = iota
	waintingInRoom = iota
	playing        = iota
)

type Client struct {
	conn  *websocket.Conn
	user  models.User
	state int
}

func (client *Client) handleMessages(hub *Hub) {

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			hub.unregister <- client
			break
		}
		client.conn.WriteMessage(websocket.TextMessage, message)
		fmt.Println(string(message))
	}
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
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
			fmt.Println("client connected ID=" + fmt.Sprint(newClient.user.ID))
			hub.clients[newClient] = true

		case leavingClient := <-hub.unregister:
			delete(hub.clients, leavingClient)
			for client := range hub.clients {
				client.conn.WriteMessage(websocket.TextMessage, DeleteUserMessage(leavingClient.user))
			}
			fmt.Println("client disconedted ID=" + fmt.Sprint(leavingClient.user.ID))
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
		return
	}
	var user models.User
	db.Db().Find(&user, userID)
	if user.ID == 0 {
		return
	}

	client := Client{conn, user, waiting}
	hub.register <- &client
	go client.handleMessages(hub)
}

func registerSocket() {

}

func RegisterRoutes(router *mux.Router) {
	hub := newHub()
	go hub.run()
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}
