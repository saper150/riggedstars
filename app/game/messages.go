package game

import (
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

type TextMessage struct {
	Type    string
	Name    string
	Payload string
}

func CreateTextMessage(name, text string) interface{} {
	return TextMessage{"text", name, text}
}
