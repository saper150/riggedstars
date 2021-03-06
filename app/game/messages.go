package game

import (
	"riggedstars/app/deck"
	"riggedstars/app/models"
	"strconv"
)

type UserMessage struct {
	Type    string
	Payload models.User
}

func NewUserMessage(user models.User) interface{} {
	return UserMessage{"newUser", user}
}

func DeleteUserMessage(user models.User) interface{} {
	return UserMessage{"deleteUser", user}
}

type SendCardsMessage struct {
	Type    string
	Payload []deck.Card
}

func CreateClientOwnCardMessage(cards []deck.Card) interface{} {
	return SendCardsMessage{"ownCards", cards}
}

func CreateSendTableCards(cards []deck.Card) interface{} {
	return SendCardsMessage{"tableCards", cards}
}

type TextMessage struct {
	Type    string
	Name    string
	Payload string
}

func CreateTextMessage(name, text string) interface{} {
	return TextMessage{"text", name, text}
}

type BetMessage struct {
	Type    string
	ID      uint
	Ammount int
}

func CreateBetMessage(client *Client, bet int) interface{} {
	return BetMessage{"bet", client.user.ID, bet}
}

type EndBetStageMessage struct {
	Type string
	Pot  int
}

func CreateEndBetStageMessage(pot int) interface{} {
	return EndBetStageMessage{"endBetStage", pot}
}

type FoldMessage struct {
	Type string
	ID   uint
}

func CreateFoldMessage(client *Client) interface{} {
	return FoldMessage{"fold", client.user.ID}
}

type ActivePlayerMessage struct {
	Type   string
	ID     uint
	Name   string
	MinBet int
}

func CreateActivePlayerMessage(client *Client, minBet int) interface{} {
	return ActivePlayerMessage{Type: "activePlayer", ID: client.user.ID, Name: client.user.Name, MinBet: minBet}
}

func CreateButtonPlayerMessage(client *Client) interface{} {
	return PlayerMessage{"button", client.user.ID, client.user.Name}
}

type PlayerMessage struct {
	Type string
	ID   uint
	Name string
}

type StartRoundInfoMessage struct {
	Type         string
	Players      map[string]PlayerInfo
	ButtonClient models.User
}

type PlayerInfo struct {
	ID    uint
	Name  string
	Stack int
}

func CreateStartRoundInfoMessage(clients map[int]*Client, gameStacks map[*Client]int, buttonClient *Client) interface{} {
	players := make(map[string]PlayerInfo)
	for seatID, client := range clients {
		players[strconv.Itoa(seatID)] = PlayerInfo{client.user.ID, client.user.Name, gameStacks[client]}
	}
	return StartRoundInfoMessage{"startRound", players, buttonClient.user}
}

type EndRoundMessage struct {
	Type    string
	Winners []PotWinner
}

type PotWinner struct {
	Winner  WinnerInfo
	Cards   []deck.Card
	Ammount int
}

type WinnerInfo struct {
	ID   uint
	Name string
}

func CreateEndRoundMessage(potWinners map[*Client]int, cards map[*Client][]deck.Card) interface{} {
	winners := make([]PotWinner, 0)
	for client, ammount := range potWinners {
		winners = append(winners, PotWinner{WinnerInfo{client.user.ID, client.user.Name}, cards[client], ammount})
	}
	return EndRoundMessage{"endRound", winners}
}
