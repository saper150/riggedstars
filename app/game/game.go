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
	blindIncrease int
	blindLevels   int
	round         *Round
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
		round:         NewRound(players, 0, 1),
	}
	for _, client := range game.clients {
		game.stacks[client] = game.startingStack
	}
	go StartRound(&game)
	return &game
}

type Round struct {
	buttonIndex  int
	roundCounter int
	roundDeck    deck.Deck
	stageBets    bets
	clientBets   bets
	folded       map[*Client]bool
	playerCards  map[*Client][]deck.Card
	clients      []*Client
}

func NewRound(players []*Client, nextButtonIndex, nextRoundCounter int) *Round {
	round := Round{
		roundDeck:    deck.ShufeledDeck(),
		playerCards:  make(map[*Client][]deck.Card),
		stageBets:    make(bets),
		clientBets:   make(bets),
		folded:       make(map[*Client]bool),
		buttonIndex:  nextButtonIndex,
		roundCounter: nextRoundCounter,
		clients:      players,
	}

	for _, client := range round.clients {
		round.stageBets[client] = 0
		round.clientBets[client] = 0
		round.folded[client] = false
		round.playerCards[client] = []deck.Card{round.roundDeck.Next(), round.roundDeck.Next()}
		client.sendMessage <- CreateClientOwnCardMessage(round.playerCards[client])
	}

	return &round
}

func (game *Game) Bet(client *Client, ammount int) {
	game.round.stageBets[client] += ammount
	game.stacks[client] -= ammount
}

func StartRound(game *Game) {

	//Small Blind
	game.Bet(game.round.clients[(game.round.buttonIndex+1)%len(game.round.clients)], game.blind)
	//Big Blind
	game.Bet(game.round.clients[(game.round.buttonIndex+2)%len(game.round.clients)], game.blind*2)

	//preflop
	activePlayerIndex := (game.round.buttonIndex + 3) % len(game.round.clients)
	betStage(game, game.round, activePlayerIndex)

	//flop
	floppedCards := game.round.roundDeck.Flop()
	game.broadcast(CreateSendTableCards(floppedCards))
	betStage(game, game.round, 1)

	//turn
	turnCard := game.round.roundDeck.TableCard()
	game.broadcast(CreateSendTableCards(turnCard))
	betStage(game, game.round, 1)

	//river
	riverCard := game.round.roundDeck.TableCard()
	game.broadcast(CreateSendTableCards(riverCard))
	betStage(game, game.round, 1)

	//showdown

	//nextRound
	//TODO: send end round message to reset state
	game.round = NewRound(game.clients, game.round.buttonIndex+1, game.round.roundCounter+1)
	StartRound(game)
}

func betStage(game *Game, round *Round, activePlayerIndex int) {
	activePlayers := activePlayersCount(round.folded)
	clientsActions := 0
	activePlayerIndex = round.nextActivePlayerIndex(activePlayerIndex)
	for !isBetStageOver(game.round.stageBets, game.round.folded, clientsActions, activePlayers) {
		maxBet := round.stageBets.maxBet()

		minBet := maxBet - round.stageBets[game.clients[activePlayerIndex]]
		game.broadcast(CreateActivePlayerMessage(round.clients[activePlayerIndex], minBet))
		bettingStageMessagesHandler(game, activePlayerIndex)
		activePlayerIndex = round.nextActivePlayerIndex((activePlayerIndex + 1) % len(round.clients))
		clientsActions++
	}
	round.clientBets.add(round.stageBets)
	round.stageBets.reset()
}

func (stageBets bets) maxBet() int {
	maxBet := 0
	for _, bet := range stageBets {
		if bet > maxBet {
			maxBet = bet
		}
	}
	return maxBet
}

func (round *Round) nextActivePlayerIndex(index int) int {
	activePlayerIndex := index
	for round.folded[round.clients[activePlayerIndex]] {
		activePlayerIndex = (activePlayerIndex + 1) % len(round.clients)
	}
	return activePlayerIndex
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

func isBetStageOver(stageBets map[*Client]int, folded map[*Client]bool, clientsActions int, activePlayers int) bool {
	return (areBetsEqual(stageBets, folded) && allActed(clientsActions, activePlayers))
}

func activePlayersCount(folded map[*Client]bool) int {
	activePlayers := 0
	for _, fold := range folded {
		if !fold {
			activePlayers++
		}
	}
	return activePlayers
}

func allActed(clientsActionsCount, activePlayersCount int) bool {
	return clientsActionsCount >= activePlayersCount
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

func bettingStageMessagesHandler(game *Game, activePlayerIndex int) {
	for {
		message := <-game.gameplayChan
		if message.From == game.round.clients[activePlayerIndex] {
			switch message.Message.Type {
			case "bet":
				betSize := int(message.Message.Payload.(float64))
				game.Bet(game.round.clients[activePlayerIndex], betSize)
				message.From.broadcast(game, CreateBetMessage(message.From, betSize))
			case "raise":
			case "fold":
				game.round.folded[game.clients[activePlayerIndex]] = true
				message.From.broadcast(game, CreateFoldMessage(message.From))
			case "check":
			}
			break
		}
	}
}

func (game *Game) addClient(client *Client) {
	game.clients = append(game.clients, client)
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
