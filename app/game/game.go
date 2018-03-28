package game

import (
	"riggedstars/app/deck"
)

type bets map[*Client]int

type Game struct {
	clients       []*Client
	buttonClient  *Client
	stacks        map[*Client]int
	gameplayChan  chan clientCommand
	blind         int
	startingStack int
}

func StartGame(clients map[*Client]bool) *Game {

	players := make([]*Client, 0)
	for player := range clients {
		players = append(players, player)
	}

	game := Game{
		clients:       players,
		buttonClient:  players[0],
		gameplayChan:  make(chan clientCommand),
		stacks:        make(map[*Client]int),
		blind:         10,
		startingStack: 500,
	}
	for _, client := range game.clients {
		game.stacks[client] = game.startingStack
	}
	go StartRound(&game)
	return &game
}

func StartRound(game *Game) {

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

	//Small Blind
	stageBets[game.clients[1%len(game.clients)]] = game.blind
	game.stacks[game.clients[1%len(game.clients)]] -= game.blind
	//Big Blind
	stageBets[game.clients[2%len(game.clients)]] = 2 * game.blind
	game.stacks[game.clients[2%len(game.clients)]] -= 2 * game.blind

	activePlayerIndex := 3 % len(game.clients)

	betStage(game, stageBets, folded, clientBets, activePlayerIndex)

	floppedCards := gameDeck.Flop()
	game.broadcast(CreateSendTableCards(floppedCards))

	//flop
	betStage(game, stageBets, folded, clientBets, 0)

	turnCard := gameDeck.TableCard()
	game.broadcast(CreateSendTableCards(turnCard))
	//turn

	betStage(game, stageBets, folded, clientBets, 0)

	riverCard := gameDeck.TableCard()
	game.broadcast(CreateSendTableCards(riverCard))

	betStage(game, stageBets, folded, clientBets, 0)

	//river
}

func betStage(game *Game, stageBets bets, folded map[*Client]bool, clientBets bets, activePlayerIndex int) {
	betStageFolded := make(map[*Client]bool)
	for client, fold := range folded {
		betStageFolded[client] = fold
	}
	clientsActions := 0
	for !isBetStageOver(stageBets, folded, betStageFolded, clientsActions) {
		for folded[game.clients[activePlayerIndex]] {
			activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		}
		maxBet := 0
		for _, bet := range stageBets {
			if bet > maxBet {
				maxBet = bet
			}
		}

		minBet := maxBet - stageBets[game.clients[activePlayerIndex]]
		game.broadcast(CreateActivePlayerMessage(game.clients[activePlayerIndex]))
		game.broadcast(CreateMinBetPlayerMessage(game.clients[activePlayerIndex], minBet))
		bettingStageMessagesHandler(game, activePlayerIndex, game.stacks, stageBets, folded)
		activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		for folded[game.clients[activePlayerIndex]] {
			activePlayerIndex = (activePlayerIndex + 1) % len(game.clients)
		}
		clientsActions++
	}
	clientBets.add(stageBets)
	stageBets.reset()
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

func isBetStageOver(stageBets map[*Client]int, folded map[*Client]bool, betStageFolded map[*Client]bool, clientsActions int) bool {
	return (areBetsEqual(stageBets, folded) && allActed(clientsActions, betStageFolded))
}

func allActed(clientsActions int, folded map[*Client]bool) bool {
	activeClients := 0
	for _, fold := range folded {
		if !fold {
			activeClients++
		}
	}
	if activeClients > clientsActions {
		return false
	}
	return true
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

func bettingStageMessagesHandler(game *Game, activePlayerIndex int, stacks map[*Client]int, stageBets bets, folded map[*Client]bool) {
	for {
		message := <-game.gameplayChan
		if message.From == game.clients[activePlayerIndex] {
			switch message.Message.Type {
			case "bet":
				betSize := int(message.Message.Payload.(float64))
				stacks[game.clients[activePlayerIndex]] -= betSize
				stageBets[game.clients[activePlayerIndex]] += betSize
				message.From.broadcast(game, CreateBetMessage(message.From, betSize))
			case "raise":
			case "fold":
				folded[game.clients[activePlayerIndex]] = true
				message.From.broadcast(game, CreateFoldMessage(message.From))
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
