package game

import (
	"fmt"
)

type clientCommand struct {
	From    *Client
	Message genericCommand
}

type Room struct {
	ID       int
	Name     string
	Clients  map[*Client]bool
	Commands chan clientCommand
	Leave    chan *Client
	Join     chan *Client
}

func (room *Room) run(hub *Hub) {
	for {
		select {
		case command := <-room.Commands:
			handleMessages(room, command)
		case client := <-room.Leave:
			delete(room.Clients, client)
			room.sendToEveryOneExcept(nil, DeleteUserMessage(client.user))
		case client := <-room.Join:
			fmt.Printf("CLient %d joined room %d", client.user.ID, room.ID)
			for c := range room.Clients {
				client.sendMessage <- NewUserMessage(c.user)
			}
			room.Clients[client] = true
			go client.handleRoom(room)
			room.sendToEveryOneExcept(client, NewUserMessage(client.user))
		}
	}
}

func (client *Client) handleRoom(room *Room) {

	go func() {
		<-client.disconect
		room.Leave <- client
	}()

	for message := range client.listen {
		room.Commands <- clientCommand{client, message}
	}
}

func handleMessages(room *Room, command clientCommand) {
	switch command.Message.Type {
	case "text":
		textMessage := CreateTextMessage(command.Message.Payload.(string))
		room.sendToEveryOneExcept(command.From, textMessage)
	}
}

func (room *Room) sendToEveryOneExcept(except *Client, message interface{}) {
	for client := range room.Clients {
		if client != except {
			client.sendMessage <- message
		}
	}
}
