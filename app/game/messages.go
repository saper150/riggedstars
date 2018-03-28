package game

import (
	"riggedstars/app/deck"
	"riggedstars/app/models"
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
	Payload string
}

func CreateTextMessage(text string) interface{} {
	return TextMessage{"text", text}
}

type BetMessage struct {
	Type    string
	ID      uint
	Ammount int
}

func CreateBetMessage(client *Client, bet int) interface{} {
	return BetMessage{"bet", client.user.ID, bet}
}

type FoldMessage struct {
	Type string
	ID   uint
}

func CreateFoldMessage(client *Client) interface{} {
	return FoldMessage{"fold", client.user.ID}
}

type ActivePlayerMessage struct {
	Type string
	ID   uint
	Name string
}

func CreateActivePlayerMessage(client *Client) interface{} {
	return ActivePlayerMessage{Type: "activePlayer", ID: client.user.ID, Name: client.user.Name}
}

func CreateMinBetPlayerMessage(client *Client, minBet int) interface{} {
	return BetMessage{"minBet", client.user.ID, minBet}
}
