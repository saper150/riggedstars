package game

import (
	"riggedstars/app/models"
	"riggedstars/app/deck"
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

type ClientOwnCardMessage struct {
	Type string
	Payload deck.Card
}

func CreateClientOwnCardMessage(card deck.Card) interface{} {
	return ClientOwnCardMessage{"ownCard", card}
}

type TextMessage struct {
	Type    string
	Payload string
}

func CreateTextMessage(text string) interface{} {
	return TextMessage{"text", text}
}
