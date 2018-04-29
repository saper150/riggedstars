package gameTests

import (
	"riggedstars/app/deck"
	"riggedstars/app/game"
)

func testHistogramHandCompareThreeOfAKindVsPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 6}, deck.Card{Value: 6}}
	name := "comparision function pair vs three of a kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}

	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed three of a kind should beat pair"
	}
}

func testHistogramHandCompareTwoPairVsPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 6}, deck.Card{Value: 8}}
	name := "comparision function pair vs two pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}

	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed two pair should beat pair"
	}
}

func testHistogramHandCompareFourOfAKindVsPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 6}, deck.Card{Value: 6}}
	name := "comparision function pair vs four of a kind"
	winners := getWinnerHand(tableCards, playerCards)

	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed four of a kind should beat pair"
	}
}

func testHistogramHandCompareFourOfAKindVsThreeOfAKind() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 6}, deck.Card{Value: 10}}
	name := "comparision function three of a kind vs Four of a kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Four of a kind should beat three of a kind"
	}
}

func testHistogramHandCompareFourOfAKindVsFullHouse() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 2}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 6}, deck.Card{Value: 6}}
	name := "comparision function full house vs four of a kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed four of a kind should beat full house"
	}
}
func testHistogramHandCompareThreeOfAKindVsFullHouse() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 11}}
	name := "comparision function full house vs three of a kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed full house should beat three of a kind"
	}
}

func testHistogramHandComparePairVsHighCard() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 12}}
	name := "comparision function pair vs high card"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed pair should beat high card"
	}
}

func testHandComparePairVsStraight() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 3, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 7}}
	name := "comparision function pair vs straight"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straight should beat pair"
	}
}

func testHandComparePairVsStraight2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 7}}
	name := "comparision function pair vs straight"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straight should beat pair"
	}
}

func testHandComparePairVsStraight3() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 5, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 1}, deck.Card{Value: 3}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 7}}
	name := "comparision function pair vs straight"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straight should beat pair"
	}
}

func testHandComparePairVsStraight4() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 1}, deck.Card{Value: 13}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 7}}
	name := "comparision function pair vs straight"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straight should beat pair"
	}
}

func testHandComparePairVsFlush() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 12, Suit: "clubs"}, deck.Card{Value: 3, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 11}}
	name := "comparision function pair vs flush"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed flush should beat pair"
	}
}

func testHandCompareStraightVsFlush() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2, Suit: "clubs"}, deck.Card{Value: 3, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 7}}
	name := "comparision function straight vs flush"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed flush should beat straight"
	}
}

func testHandCompareFlushVsStraightFlush() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "clubs"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 2, Suit: "spades"}, deck.Card{Value: 3, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5, Suit: "clubs"}, deck.Card{Value: 3, Suit: "clubs"}}
	name := "comparision function flush vs straightflush"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straightflush should beat flush"
	}
}

func testHandCompareStraightVsStraightFlush() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 7, Suit: "clubs"}, deck.Card{Value: 10, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 7}, deck.Card{Value: 10}}
	name := "comparision function pair vs straightflush"
	winners := getWinnerHand(tableCards, playerCards)

	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed straightflush should beat pair"
	}
}
