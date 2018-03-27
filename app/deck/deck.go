package deck

import "math/rand"

type Card struct {
	Suit  string
	Value int
}

type Deck []Card

func (deck *Deck) Next() Card {
	var card Card
	card, *deck = (*deck)[0], (*deck)[1:]
	return card
}

func (deck *Deck) Flop() []Card {
	deck.Next()
	return []Card{deck.Next(), deck.Next(), deck.Next()}
}

func ShufeledDeck() Deck {

	deck := make(Deck, 52)

	for i, suit := range []string{"clubs", "diamonds", "hearts", "spades"} {
		for j := 0; j < 13; j++ {
			deck[i*13+j].Suit = suit
			deck[i*13+j].Value = j
		}
	}

	for i := 51; i > 0; i-- {
		newIndex := rand.Intn(i)
		deck[i], deck[newIndex] = deck[newIndex], deck[i]
	}
	return deck
}
