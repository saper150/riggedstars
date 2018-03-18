package game

import (
	"encoding/json"
	"riggedstars/app/deck"
	"riggedstars/app/models"
)

type UserMessage struct {
	Type    string
	Payload models.User
}

func NewUserMessage(user models.User) []byte {
	js, _ := json.Marshal(UserMessage{"newUser", user})
	return js
}

func DeleteUserMessage(user models.User) []byte {
	js, _ := json.Marshal(UserMessage{"deleteUser", user})
	return js
}

type TextMessage struct {
	Type    string
	Payload string
}

func CreateTextMessage(text string) []byte {
	js, _ := json.Marshal(TextMessage{"text", text})
	return js
}

func gameStartMessage() []byte {
	js, _ := json.Marshal(TextMessage{"gameStart", ""})
	return js
}

type DealCard struct {
	Type    string
	Payload deck.Card
}

func dealCardMessage(card deck.Card) []byte {
	js, _ := json.Marshal(DealCard{"dealCard", card})
	return js
}
