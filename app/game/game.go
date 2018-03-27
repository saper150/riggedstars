package game

import (
	"riggedstars/app/deck"
)

type bets map[*Client]int

type Game struct {
	clients      []*Client
	buttonClient *Client
	stacks       map[*Client]int
	gameplayChan chan clientCommand
}

func StartGame(clients map[*Client]bool) *Game {

	players := make([]*Client, 0)
	for player := range clients {
		players = append(players, player)
	}

	game := Game{
		clients:      players,
		buttonClient: players[0],
		gameplayChan: make(chan clientCommand),
		stacks:       make(map[*Client]int),
	}
	for _, client := range game.clients {
		game.stacks[client] = 500
	}
	go StartRound(&game)
	return &game
}

func StartRound(game *Game) {
	//cards
	gameDeck := deck.ShufeledDeck()
	stageBets := make(bets)
	clientBets := make(bets)

	folded := make(map[*Client]bool)
	playerCards := make(map[*Client][]deck.Card)
	for _, client := range game.clients {
		stageBets[client] = 0
		clientBets[client] = 0
		folded[client] = false
		playerCards[client] = []deck.Card{gameDeck.Next(), gameDeck.Next()}
		client.sendMessage <- CreateClientOwnCardMessage(playerCards[client])
	}

	stageBets[game.clients[1%len(game.clients)]] = 10
	game.stacks[game.clients[1%len(game.clients)]] -= 10

	stageBets[game.clients[2%len(game.clients)]] = 20
	game.stacks[game.clients[2%len(game.clients)]] -= 20

	activePlayerIndex := 3 % len(game.clients)

	for !areBetsEqual(stageBets, folded) {
		bettingStageMessagesHandler(game, activePlayerIndex, game.stacks, stageBets)
		activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		for folded[game.clients[activePlayerIndex]] {
			activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		}
	}
	clientBets.add(stageBets)
	stageBets.reset()

	floppedCards := gameDeck.Flop()
	game.broadcast(CreateSendTableCards(floppedCards))

	//flop
	for !areBetsEqual(stageBets, folded) {
		bettingStageMessagesHandler(game, activePlayerIndex, game.stacks, stageBets)
		activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		for folded[game.clients[activePlayerIndex]] {
			activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		}
	}
	clientBets.add(stageBets)
	stageBets.reset()

	gameDeck.Next() // burn
	turnCard := []deck.Card{gameDeck.Next()}
	game.broadcast(CreateSendTableCards(turnCard))
	//turn

	for !areBetsEqual(stageBets, folded) {
		bettingStageMessagesHandler(game, activePlayerIndex, game.stacks, stageBets)
		activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		for folded[game.clients[activePlayerIndex]] {
			activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		}
	}

	clientBets.add(stageBets)
	stageBets.reset()

	gameDeck.Next() // burn
	riverCard := []deck.Card{gameDeck.Next()}
	game.broadcast(CreateSendTableCards(riverCard))
	//river
}

func (bets bets) add(from map[*Client]int) {
	for client, bet := range bets {
		bets[client] = bet + from[client]
	}
}

func (bets bets) reset() {
	for client := range bets {
		bets[client] = 0
	}
}

func areBetsEqual(bets map[*Client]int, folded map[*Client]bool) bool {
	set := make(map[int]interface{})
	for client, bet := range bets {
		if !folded[client] {
			set[bet] = ""
			if len(set) != 1 {
				return false
			}
		}
	}
	return true
}

func bettingStageMessagesHandler(game *Game, activePlayerIndex int, stacks map[*Client]int, stageBets bets) {
	for {
		message := <-game.gameplayChan
		if message.From == game.clients[activePlayerIndex] {
			switch message.Message.Type {
			case "bet":
				betSize := int(message.Message.Payload.(float64))
				stacks[game.clients[activePlayerIndex]] -= betSize
				stageBets[game.clients[activePlayerIndex]] += betSize
				message.From.broadcast(game, CreateBetMessage(message.From, betSize))
			case "rase":
			case "fold":
			case "check":
			}
			break
		}
	}
}

func (game *Game) broadcast(message interface{}) {
	for _, client := range game.clients {
		client.sendMessage <- message
	}
}

func (me *Client) broadcast(game *Game, message interface{}) {
	for _, client := range game.clients {
		if me != client {
			client.sendMessage <- message
		}
	}
}

func getWinner(playerCards map[*Client][]deck.Card) *Client {
	for client := range playerCards {
		return client
	}
	return nil
}
