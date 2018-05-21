package game

import (
	"riggedstars/app/db"
	"riggedstars/app/deck"
	"time"
)

type bets map[*Client]int

type Game struct {
	clients       map[int]*Client
	clientsState  map[*Client]bool
	buttonClient  *Client
	stacks        map[*Client]int
	gameplayChan  chan clientCommand
	blind         int
	startingStack int
	blindIncrease int
	blindLevels   int
	round         *Round
	maxPlayers    int
}

func StartGame(clients map[*Client]bool, maxClients int) *Game {
	clientsArray := make([]*Client, 0)
	for client := range clients {
		clientsArray = append(clientsArray, client)
	}
	players := make(map[int]*Client)
	for i := 0; i < len(clientsArray); i++ {
		players[i] = clientsArray[i]
	}
	game := Game{
		clients:       players,
		clientsState:  make(map[*Client]bool),
		gameplayChan:  make(chan clientCommand),
		stacks:        make(map[*Client]int),
		blind:         10,
		startingStack: 500,
		round:         NewRound(players, 0, 1, maxClients),
		maxPlayers:    maxClients,
	}
	for _, client := range game.clients {
		game.clientsState[client] = true
		game.stacks[client] = game.startingStack
		db.ChangeStack(client.user, -game.startingStack)
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
	clients      map[int]*Client
	maxPlayers   int
	tableCards   []deck.Card
	allInPlayers map[*Client]bool
}

func NewRound(players map[int]*Client, nextButtonIndex, nextRoundCounter, maxClients int) *Round {
	round := Round{
		roundDeck:    deck.ShufeledDeck(),
		playerCards:  make(map[*Client][]deck.Card),
		stageBets:    make(bets),
		clientBets:   make(bets),
		folded:       make(map[*Client]bool),
		buttonIndex:  nextButtonIndex,
		roundCounter: nextRoundCounter,
		clients:      players,
		maxPlayers:   maxClients,
		tableCards:   make([]deck.Card, 0),
		allInPlayers: make(map[*Client]bool),
	}

	for _, client := range round.clients {
		round.stageBets[client] = 0
		round.clientBets[client] = 0
		round.folded[client] = false
		round.allInPlayers[client] = false
		round.playerCards[client] = []deck.Card{round.roundDeck.Next(), round.roundDeck.Next()}
		client.sendMessage <- CreateClientOwnCardMessage(round.playerCards[client])
	}

	return &round
}

func (round *Round) nextPlayerIndex(index int) int {
	nextPlayerIndex := (index + 1) % round.maxPlayers
	for round.clients[nextPlayerIndex] == nil {
		nextPlayerIndex = (nextPlayerIndex + 1) % round.maxPlayers
	}
	return nextPlayerIndex
}

func (game *Game) Bet(client *Client, ammount int) {
	game.round.stageBets[client] += ammount
	game.stacks[client] -= ammount
}

func (game *Game) BetAllIn(client *Client, ammount int) {
	game.Bet(client, ammount)
	game.round.allInPlayers[client] = true
}

func (game *Game) Blind(client *Client, multi int) {
	blindSize := game.blind * multi
	if game.stacks[client] <= blindSize {
		game.BetAllIn(client, game.stacks[client])
		game.broadcast(CreateBetMessage(client, game.stacks[client]))
	} else {
		game.Bet(client, game.blind*multi)
		game.broadcast(CreateBetMessage(client, game.blind*multi))
	}
}

func StartRound(game *Game) {

	game.broadcast(CreateStartRoundInfoMessage(game.round.clients, game.stacks, game.round.clients[game.round.buttonIndex]))

	smallBlindSeatIndex := game.round.nextPlayerIndex(game.round.buttonIndex)
	game.Blind(game.round.clients[smallBlindSeatIndex], 1)
	bigBlindSeatIndex := game.round.nextPlayerIndex(smallBlindSeatIndex)
	game.Blind(game.round.clients[bigBlindSeatIndex], 2)

	//preflop
	betStage(game, game.round.nextPlayerIndex(bigBlindSeatIndex))

	//flop
	if game.round.activePlayersCount() > 1 {
		floppedCards := game.round.roundDeck.Flop()
		game.round.tableCards = append(game.round.tableCards, floppedCards...)
		game.broadcast(CreateSendTableCards(floppedCards))
		betStage(game, 1)
	}

	//turn
	if game.round.activePlayersCount() > 1 {
		turnCard := game.round.roundDeck.TableCard()
		game.round.tableCards = append(game.round.tableCards, turnCard...)
		game.broadcast(CreateSendTableCards(turnCard))
		betStage(game, 1)
	}

	//river
	if game.round.activePlayersCount() > 1 {
		riverCard := game.round.roundDeck.TableCard()
		game.round.tableCards = append(game.round.tableCards, riverCard...)
		game.broadcast(CreateSendTableCards(riverCard))
		betStage(game, 1)
	}
	//showdown
	winners := make([]*Client, 0)
	if game.round.activePlayersCount() > 1 {
		showdownCards := make(map[*Client][]deck.Card)
		for client, fold := range game.round.folded {
			if !fold {
				showdownCards[client] = game.round.playerCards[client]
			}
		}
		winners = GetWinnerHand(game.round.tableCards, showdownCards)
	} else {
		for client, fold := range game.round.folded {
			if !fold {
				winners = append(winners, client)
			}
		}
	}
	//TODO: all ins
	pot := game.round.clientBets.sum()
	potWinners := make(map[(*Client)]int)
	for _, client := range winners {
		ammount := pot / len(winners)
		potWinners[client] = ammount
		game.stacks[client] += ammount
	}

	//nextRound
	game.broadcast(CreateEndRoundMessage(potWinners))

	game.deleteDisconnectedClients()

	//TODO: evade using sleep
	time.Sleep(time.Second * 3)

	for game.round.countPlayers() <= 1 {
		time.Sleep(time.Second * 1)
	}
	game.round = NewRound(game.clients, game.nextPlayerIndex(game.round.buttonIndex), game.round.roundCounter+1, game.maxPlayers)
	//TODO: find better solution lol
	StartRound(game)
}

func (round *Round) countPlayers() int {
	sum := 0
	for range round.clients {
		sum++
	}
	return sum
}

func betStage(game *Game, activePlayerIndex int) {
	clientsActions := 0
	activePlayerIndex = game.round.nextActivePlayerIndex(activePlayerIndex)
	activePlayersOnStart := game.round.activePlayersCount()
	for !isBetStageOver(game.round, clientsActions, activePlayersOnStart) {
		if !game.clientsState[game.round.clients[activePlayerIndex]] {
			game.fold(game.round.clients[activePlayerIndex])
		} else {
			maxBet := game.round.maxBet()
			minBet := maxBet - game.round.stageBets[game.round.clients[activePlayerIndex]]
			game.broadcast(CreateActivePlayerMessage(game.round.clients[activePlayerIndex], minBet))
			bettingStageMessagesHandler(game, activePlayerIndex, minBet)
		}
		activePlayerIndex = game.round.nextActivePlayerIndex((activePlayerIndex + 1) % len(game.round.clients))
		clientsActions++
	}
	game.round.clientBets.add(game.round.stageBets)
	game.broadcast(CreateEndBetStageMessage(game.round.stageBets.sum()))
	game.round.stageBets.reset()
}

func (game *Game) nextPlayerIndex(index int) int {
	nextPlayerIndex := (index + 1) % len(game.clients)
	for game.clients[nextPlayerIndex] == nil {
		nextPlayerIndex = (nextPlayerIndex + 1) % len(game.clients)
	}
	return nextPlayerIndex
}

func (round *Round) maxBet() int {
	maxBet := 0
	for _, bet := range round.stageBets {
		if bet > maxBet {
			maxBet = bet
		}
	}
	return maxBet
}

func (round *Round) nextActivePlayerIndex(index int) int {
	activePlayerIndex := index
	for round.folded[round.clients[activePlayerIndex]] || round.allInPlayers[round.clients[activePlayerIndex]] {
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

func (bets bets) sum() int {
	betSum := 0
	for _, bet := range bets {
		betSum += bet
	}
	return betSum
}

func isBetStageOver(round *Round, clientsActions, activePlayersOnStart int) bool {
	return round.areBetsEqual() && (clientsActions >= activePlayersOnStart || round.activePlayersCount() == 1)
}

func (round *Round) activePlayersCount() int {
	activePlayers := 0
	for _, fold := range round.folded {
		if !fold {
			activePlayers++
		}
	}
	return activePlayers
}

func (round *Round) areBetsEqual() bool {
	set := make(map[int]interface{})
	for client, bet := range round.stageBets {
		if !round.folded[client] && !round.allInPlayers[client] {
			set[bet] = ""
			if len(set) != 1 {
				return false
			}
		}
	}
	return true
}

func bettingStageMessagesHandler(game *Game, activePlayerIndex int, minBet int) {
	for {
		message := <-game.gameplayChan
		if message.From == game.round.clients[activePlayerIndex] {
			switch message.Message.Type {
			case "bet":
				betSize := int(message.Message.Payload.(float64))
				if game.stacks[message.From]-betSize <= 0 {
					game.BetAllIn(message.From, betSize)
					message.From.broadcast(game, CreateBetMessage(message.From, betSize))
				} else {
					game.Bet(message.From, betSize)
					message.From.broadcast(game, CreateBetMessage(message.From, betSize))
				}
			case "raise":
			case "fold":
				game.fold(message.From)
			case "check":
			}
			break
		}
	}
}

func (game *Game) fold(client *Client) {
	game.round.folded[client] = true
	client.broadcast(game, CreateFoldMessage(client))
}

func (game *Game) addClient(client *Client) {
	index := 0
	for game.clients[index] != nil {
		index++
	}
	game.clients[index] = client
	game.stacks[client] = game.startingStack
	game.clientsState[client] = true
	game.round.folded[client] = true
	db.ChangeStack(client.user, -game.startingStack)
}

func (game *Game) deleteClient(client *Client) {
	game.clientsState[client] = false
}

func (game *Game) deleteDisconnectedClients() {
	for seatID, client := range game.clients {
		if !game.clientsState[client] {
			delete(game.clients, seatID)
		}
	}
}

func (game *Game) broadcast(message interface{}) {
	for _, client := range game.clients {
		if game.clientsState[client] {
			client.sendMessage <- message
		}
	}
}

func (me *Client) broadcast(game *Game, message interface{}) {
	for _, client := range game.clients {
		if me != client && game.clientsState[client] {
			client.sendMessage <- message
		}
	}
}
