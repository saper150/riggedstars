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

func getWinnerHand(tableCards []deck.Card, clientCards map[*Client][]deck.Card) []*Client {

	fullPlayerCards := make(map[*Client][]deck.Card)
	bestTypeSoFar := make(map[*Client]int)
	for client, cards := range clientCards {
		fullPlayerCards[client] = append(tableCards, cards...)
		bestTypeSoFar[client] = groupByValue(fullPlayerCards[client]).getHistogramHand()
		straight, _ := isStraight(fullPlayerCards[client])
		if straight {
			if bestTypeSoFar[client] < Straight {
				bestTypeSoFar[client] = Straight
			}
		}

		flush, _ := isFlush(fullPlayerCards[client])
		if flush {
			if bestTypeSoFar[client] < Flush {
				bestTypeSoFar[client] = Flush
			}
		}

		if straightFlush, _ := isStraightFlush(fullPlayerCards[client]); straightFlush {
			bestTypeSoFar[client] = StraightFlush
		}
	}

	tiedTypePlayers, maxType := bestHandType(bestTypeSoFar, fullPlayerCards)
	winners := resolveTie(tiedTypePlayers, maxType)
	winnersArr := make([]*Client, 0)

	for client, _ := range winners {
		winnersArr = append(winnersArr, client)
	}

	return winnersArr

}

func resolveTie(tiedPlayersCards map[*Client][]deck.Card, maxType int) map[*Client][]deck.Card {
	switch maxType {
	case HighCard:
		return resolveHighCardTie(tiedPlayersCards, 5)
	case Pair:
		return resolvePairTie(tiedPlayersCards)
	default:
		return tiedPlayersCards
	}
}

func resolveHighCardTie(playerCards map[*Client][]deck.Card, cardsCount int) map[*Client][]deck.Card {

	for _, cards := range playerCards {

		sort.Slice(cards, func(i, j int) bool {
			if cards[i].Value == 1 {
				return true
			}
			if cards[j].Value == 1 {
				return false
			}
			return cards[i].Value > cards[j].Value
		})
	}
	for i := 0; i < cardsCount; i++ {
		maxCard := 0
		for _, cards := range playerCards {
			if maxCard == 1 || cards[i].Value == 1 {
				maxCard = 1
			} else if cards[i].Value > maxCard {
				maxCard = cards[i].Value
			}
		}
		tempPlayerCards := make(map[*Client][]deck.Card)
		for client, cards := range playerCards {
			if cards[i].Value == maxCard {
				tempPlayerCards[client] = cards
			}
		}
		playerCards = tempPlayerCards
	}

	return playerCards
}

func resolvePairTie(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	highestPair := make(map[*Client]int)
	high := 0
	for client, cards := range playerCards {
		clientHighestPair, _ := getHighestPair(cards)
		if clientHighestPair > high || clientHighestPair == 1 {
			high = clientHighestPair
		}
		highestPair[client] = clientHighestPair
	}

	tempPlayerCards := make(map[*Client][]deck.Card)
	for client, cards := range playerCards {
		if highestPair[client] == high {
			tempCards := make([]deck.Card, 0)
			for _, card := range cards {
				if card.Value != high {
					tempCards = append(tempCards, card)
				}
			}
			tempPlayerCards[client] = tempCards
		}
	}
	playerCards = tempPlayerCards

	return resolveHighCardTie(playerCards, 3)
}

func getHighestPair(cards []deck.Card) (int, []deck.Card) {
	histogramPre := make(histogram)
	for _, card := range cards {
		histogramPre[card.Value] = histogramPre[card.Value] + 1
	}
	maxPairValue := 0
	for value, count := range histogramPre {
		if count != 2 {
			continue
		}
		if value > maxPairValue || value == 1 {
			maxPairValue = value
		}
	}

	restCards := make([]deck.Card, 0)
	for _, card := range cards {
		if card.Value != maxPairValue {
			restCards = append(restCards, card)
		}
	}

	return maxPairValue, restCards
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

func groupByValue(cards []deck.Card) histogram {
	histogramPre := make(histogram)
	for _, card := range cards {
		histogramPre[card.Value] = histogramPre[card.Value] + 1
	}

	histogram := make(histogram)
	for _, count := range histogramPre {
		histogram[count] = histogram[count] + 1
	}
	return histogram
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

func bestHandType(bestHands map[*Client]int, cards map[*Client][]deck.Card) (map[*Client][]deck.Card, int) {

	tiedPlayers := make(map[*Client][]deck.Card)
	maxType := -1
	for _, handType := range bestHands {
		if handType > maxType {
			maxType = handType
		}
	}

	for client, handType := range bestHands {
		if handType == maxType {
			tiedPlayers[client] = cards[client]
		}
	}

	return tiedPlayers, maxType
}
