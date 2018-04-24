package game

import (
	"riggedstars/app/deck"
	"sort"
)

const (
	HighCard      = iota
	Pair          = iota
	TwoPairs      = iota
	ThreeOfAKind  = iota
	Straight      = iota
	Flush         = iota
	FullHouse     = iota
	FourOfAKind   = iota
	StraightFlush = iota
	Poker         = iota
)

type histogram map[int]int

func getWinnerHand(tableCards []deck.Card, clientCards map[*Client][]deck.Card) *Client {

	bestTypeSoFar := make(map[*Client]int)
	for client, cards := range clientCards {
		fullPlayerCards := append(tableCards, cards...)
		bestTypeSoFar[client] = groupByValue(fullPlayerCards).getHistogramHand()
		straight, _ := isStraight(fullPlayerCards)
		if straight {
			if bestTypeSoFar[client] < Straight {
				bestTypeSoFar[client] = Straight
			}
		}

		flush, _ := isFlush(fullPlayerCards)
		if flush {
			if bestTypeSoFar[client] < Flush {
				bestTypeSoFar[client] = Flush
			}
		}

		if straightFlush, _ := isStraightFlush(fullPlayerCards); straightFlush {
			bestTypeSoFar[client] = StraightFlush
		}

	}

	return bestHandType(bestTypeSoFar)
}

func groupByValue(cards []deck.Card) histogram {
	histogramPre := make(histogram)
	for _, card := range cards {
		histogramPre[card.Value] = histogramPre[card.Value] + 1
	}
	histogram := make(histogram)
	for _, count := range histogramPre {
		histogram[count] = histogramPre[count] + 1
	}
	return histogram
}

func isStraight(cards []deck.Card) (bool, deck.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})
	maxCard := deck.Card{}
	sortedCount := 0
	for i := 1; i < len(cards); i++ {
		if cards[i-1].Value == cards[i].Value-1 {
			sortedCount++
		} else {
			sortedCount = 0
		}
		if sortedCount >= 5 {
			maxCard = cards[i]
		}
	}

	return sortedCount >= 5, maxCard
}

func isFlush(cards []deck.Card) (bool, []deck.Card) {
	suits := make(map[string][]deck.Card)

	suits["clubs"] = make([]deck.Card, 0)
	suits["spades"] = make([]deck.Card, 0)
	suits["hearts"] = make([]deck.Card, 0)
	suits["diamonds"] = make([]deck.Card, 0)

	for _, card := range cards {
		suits[card.Suit] = append(suits[card.Suit], card)
	}

	for _, suitedCards := range suits {
		if len(suitedCards) >= 5 {
			return true, suitedCards
		}

	}
	return false, nil
}

func isStraightFlush(cards []deck.Card) (bool, deck.Card) {
	if isFlush, flushCards := isFlush(cards); isFlush {
		if straight, highCard := isStraight(flushCards); straight {
			return true, highCard
		}
	}
	return false, deck.Card{}
}

func (histogram histogram) getHistogramHand() int {
	if _, ok := histogram[4]; ok {
		return FourOfAKind
	}
	if _, ok := histogram[3]; ok {
		if _, ok := histogram[2]; ok {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if val, ok := histogram[2]; ok && val >= 2 {
		return TwoPairs
	}
	if _, ok := histogram[2]; ok {
		return Pair
	}
	return HighCard
}

func (data histogram) isTwoPair() bool {
	counter := 0
	for _, count := range data {
		if count == 2 {
			counter++
		}
	}
	return counter >= 2
}

func bestHandType(bestHands map[*Client]int) *Client {

	var winningPlayer *Client
	maxType := -1
	for client, handType := range bestHands {
		if handType > maxType {
			maxType = handType
			winningPlayer = client
		}
	}
	return winningPlayer
}
