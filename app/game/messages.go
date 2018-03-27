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
	ammount int
}

func CreateBetMessage(client *Client, bet int) interface{} {
	return BetMessage{"bet", client.user.ID, bet}
}
