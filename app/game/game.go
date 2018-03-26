package game

import (
	"riggedstars/app/deck"
)

type Game struct{
	clients []*Client
	buttonClient *Client
	stacks map[*Client] int
	cards map[*Client] [2]deck.Card
	gameplayChan chan clientCommand
}

func StartGame(clients []*Client){
	game := Game{clients: clients, 
		buttonClient: clients[0],
		 gameplayChan: make(chan clientCommand),
		}
	for _, client := range clients{
		game.stacks[client] = 500
	}

}

func StartRound(game *Game){
	//cards
	gameDeck := deck.ShufeledDeck()
	var bets map[*Client]int
	var folded map[*Client]bool

		for _, client := range game.clients {
			bets[client] = 0
			folded[client] = false
			game.cards[client] = [2]deck.Card{gameDeck.Next(),gameDeck.Next()}
			client.sendMessage <- CreateClientOwnCardMessage(game.cards[client][0])
			client.sendMessage <- CreateClientOwnCardMessage(game.cards[client][1])
		}

		bets[game.clients[1 % len(game.clients)]] = 10
		game.clients[1 % len(game.clients)]-=10

		bets[game.clients[2 % len(game.clients)]] = 20
		game.clients[2 % len(game.clients)]-= 20

	var activePlayer *Client
	activePlayer = game.clients[3 % len(game.clients)]
	activePlayerIndex := 3 % len(game.clients)

	for !areBetsEqual() {
		for {
			message := <- game.gameplayChan
			if message.From == activePlayer {
				switch message.Message.Type {
				case "bet":
				case "rase":
				case "fold":
				case "check":
				}
				break
			}
		}
		activePlayerIndex = (activePlayer + 1) % len(game.clients)
	}

	//preflop

	//flop

	//turn

	//river
}


func areBetsEqual(bets map[*Client]int, folded map[*Client]bool) {
	var set map[int]interface{}
	for bet, client := range bets {
		if !folded[client]{
			set[bet] = struct{}
			if len(set) != 1 {
				return false
			 }
		}
	}
	return true
}

