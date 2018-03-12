package game

import (
	"encoding/json"
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
