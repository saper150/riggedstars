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

func GetWinnerHand(tableCards []deck.Card, clientCards map[*Client][]deck.Card) []*Client {

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
		return resolvePairVsPair(tiedPlayersCards)
	case TwoPairs:
		return resolveTwoPairVsTwoPair(tiedPlayersCards)
	case ThreeOfAKind:
		return resolveThreeOfAkindVsThreeOfAKind(tiedPlayersCards)
	case FourOfAKind:
		return resolveFourOfAKindVsFourOfAKind(tiedPlayersCards)
	case FullHouse:
		return resolveFullHouseVsFullHouse(tiedPlayersCards)
	case Straight:
		return resolveStraightVsStraight(tiedPlayersCards)
	case Flush:
		return resolveFlushVsFlush(tiedPlayersCards)
	case StraightFlush:
		return resolveStraightVsStraight(tiedPlayersCards)
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

func resolvePairVsPair(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	return resolveHighCardTie(resolveGroupTypeTie(playerCards, 2), 3)
}

func resolveGroupTypeTie(playerCards map[*Client][]deck.Card, group int) map[*Client][]deck.Card {
	highestGroup := make(map[*Client]int)
	restPlayerCards := make(map[*Client][]deck.Card)
	highestGroupValue := 0
	for client, cards := range playerCards {
		clientHighestPair, restCards := getHighestGroup(cards, group)
		restPlayerCards[client] = restCards
		if highestGroupValue != 1 && clientHighestPair > highestGroupValue || clientHighestPair == 1 {
			highestGroupValue = clientHighestPair
		}
		highestGroup[client] = clientHighestPair
	}
	for client, _ := range restPlayerCards {
		if highestGroup[client] != highestGroupValue {
			delete(restPlayerCards, client)
		}
	}
	return restPlayerCards
}

func resolveTwoPairVsTwoPair(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	return resolveHighCardTie(resolveGroupTypeTie(resolveGroupTypeTie(playerCards, 2), 2), 1)
}

func resolveThreeOfAkindVsThreeOfAKind(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	return resolveHighCardTie(resolveGroupTypeTie(playerCards, 3), 2)
}

func resolveFourOfAKindVsFourOfAKind(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	return resolveHighCardTie(resolveGroupTypeTie(playerCards, 4), 1)
}

func resolveFullHouseVsFullHouse(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	return resolveGroupTypeTie(resolveGroupTypeTie(playerCards, 3), 2)
}

func resolveStraightVsStraight(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	maxClientStraighValue := make(map[*Client]int)
	maxStraightCardValue := 0
	for client, cards := range playerCards {
		_, maxCard := isStraight(cards)
		maxClientStraighValue[client] = maxCard.Value
		if maxStraightCardValue != 1 && maxCard.Value > maxStraightCardValue || maxCard.Value == 1 {
			maxStraightCardValue = maxCard.Value
		}
	}
	winningStraightCards := make(map[*Client][]deck.Card)
	for client, cards := range playerCards {
		if maxClientStraighValue[client] == maxStraightCardValue {
			winningStraightCards[client] = cards
		}
	}
	return winningStraightCards
}

func resolveFlushVsFlush(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {
	maxClientFlushCard := make(map[*Client]deck.Card)
	maxFlushCardValue := 0
	for client, cards := range playerCards {
		_, suitedCards := isFlush(cards)
		for _, card := range suitedCards {
			if card.Value == 1 {
				maxClientFlushCard[client] = card
				maxFlushCardValue = card.Value
				break
			}
			if card.Value > maxClientFlushCard[client].Value {
				maxClientFlushCard[client] = card
			}
			if card.Value > maxFlushCardValue && maxFlushCardValue != 1 {
				maxFlushCardValue = card.Value
			}
		}
	}
	winningFlushCards := make(map[*Client][]deck.Card)
	for client, cards := range playerCards {
		if maxClientFlushCard[client].Value == maxFlushCardValue {
			winningFlushCards[client] = cards
		}
	}

	return winningFlushCards
}

func resolveStraightFlushVsStraightFlush(playerCards map[*Client][]deck.Card) map[*Client][]deck.Card {

	return playerCards
}

func getHighestGroup(cards []deck.Card, group int) (int, []deck.Card) {
	histogramPre := make(histogram)
	for _, card := range cards {
		histogramPre[card.Value] = histogramPre[card.Value] + 1
	}
	maxPairValue := 0
	for value, count := range histogramPre {
		if count == group {
			if maxPairValue != 1 && value > maxPairValue || value == 1 {
				maxPairValue = value
			}
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
	for _, card := range cards {
		if card.Value == 1 {
			cards = append(cards, card)
			break
		}
	}
	maxCard := deck.Card{}
	hasStraight := false
	sortedCount := 1
	for i := 1; i < len(cards); i++ {
		if cards[i-1].Value == cards[i].Value {
			continue
		}
		if cards[i-1].Value == cards[i].Value-1 || cards[i-1].Value == cards[i].Value+12 {
			sortedCount++
		} else {
			sortedCount = 1
		}
		if sortedCount >= 5 {
			hasStraight = true
			maxCard = cards[i]
		}
	}
	return hasStraight, maxCard
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
