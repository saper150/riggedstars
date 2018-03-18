package game

import "riggedstars/app/deck"

type Game struct {
	players []*RoomClient
	dealer  *RoomClient
	deck    deck.Deck
}

func NewGame(room *Room) *Game {
	game := Game{make([]*RoomClient, len(room.Clients)), nil, deck.ShufeledDeck()}
	i := 0
	for player := range room.Clients {
		game.players[i] = player
		player.Client.send <- gameStartMessage()
		player.Client.send <- dealCardMessage(game.deck.Next())
		player.Client.send <- dealCardMessage(game.deck.Next())
		i++
	}

	game.dealer = game.players[0]

	return &game
}
