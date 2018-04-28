package game

import (
	"fmt"
	"riggedstars/app/deck"
)

func RunTests() {
	tests := []func() (string, bool, string){
		testHistogramHandCompareThreeOfAKindVsFullHouse,
		testHistogramHandCompareThreeOfAKindVsPair,
		testHistogramHandCompareFourOfAKindVsPair,
		testHistogramHandCompareFourOfAKindVsThreeOfAKind,
		testHistogramHandCompareFourOfAKindVsFullHouse,
		testHandComparePairVsStraight,
		testHandComparePairVsFlush,
		testHandCompareStraightVsFlush,
		testHandCompareFlushVsStraightFlush,
		testHandCompareStraightVsStraightFlush,
		testHistogramHandCompareTwoPairVsPair,
		testHandCompareHighCardvsHighCardTier1,
		testHandCompareHighCardvsHighCardTier2,
		testHandCompareHighCardvsHighCardTier3,
		testHandCompareHighCardvsHighCardTier4,
		testHandCompareHighCardvsHighCardTier5,
		testHandCompareHighCardvsHighCardTie,
		testHandCompareHighCardvsHighCardAce1,
		testHandComparePairVsPair,
		testHandComparePairVsPair2,
		testHandComparePairVsPairTie,
	}

	testCounter := 0
	for _, test := range tests {
		name, passed, result := test()
		if passed == false {
			fmt.Printf("Name: %s  Result: ", name)
			fmt.Println(result)
		} else {
			testCounter++
		}
	}
	fmt.Printf("Tests finished with %d/%d passed\n", testCounter, len(tests))
}

func testHistogramHandCompareThreeOfAKindVsPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "hearts"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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

func testHandComparePairVsFlush() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 2, Suit: "clubs"},
		deck.Card{Value: 4, Suit: "clubs"},
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 8, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

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

func testHandCompareHighCardvsHighCardTier1() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "clubs"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 1, Suit: "clubs"}, deck.Card{Value: 5, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tier 1"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card Ace should be high card 5"
	}
}

func testHandCompareHighCardvsHighCardTier2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "clubs"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 13, Suit: "clubs"}, deck.Card{Value: 5, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tier 2"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card King should be high card 5"
	}
}

func testHandCompareHighCardvsHighCardTier3() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 12, Suit: "clubs"}, deck.Card{Value: 5, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tier 3"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card Queen should be high card 5"
	}
}

func testHandCompareHighCardvsHighCardTier4() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 6, Suit: "clubs"}, deck.Card{Value: 5, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 11}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tier 4"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card Jack should be high card 5"
	}
}

func testHandCompareHighCardvsHighCardTier5() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 5, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 1, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 7, Suit: "clubs"}, deck.Card{Value: 11, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 11}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tier 5"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card 7 should be high card 5"
	}
}

func testHandCompareHighCardvsHighCardTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 8, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "clubs"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 9, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 5, Suit: "diamonds"}, deck.Card{Value: 3, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 5}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card tie"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) == 2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed expected a tie"
	}
}

func testHandCompareHighCardvsHighCardAce1() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 6, Suit: "clubs"}, deck.Card{Value: 5, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 1}, deck.Card{Value: 2}}
	name := "comparision function high card vs high card with Ace"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player2 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed high card Ace should be high card 6"
	}
}

func testHandComparePairVsPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 11, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function pair vs pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed pair of 10 with high card Jack should beat pair of 10 with high card 7"
	}
}

func testHandComparePairVsPair2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function pair vs pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed pair of 10 with high card Ace should beat pair of 10 with high card 7"
	}
}

func testHandComparePairVsPairTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*Client][]deck.Card)

	player1 := &Client{}
	player2 := &Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 4, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function pair vs pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed pair of 10 high cards: 13,12,7 should tie with pair of 10 high cards: 13,12,7"
	}
}

// func testHandCompareTwoPairsToTwoPairs() {
// 	tableCards := []deck.Card {
// 		deck.Card{Value: 2, Suit: "clubs"},
// 		deck.Card{Value: 4, Suit: "clubs"},
// 		deck.Card{Value: 6, Suit: "spades"},
// 		deck.Card{Value: 8, Suit: "hearts"},
// 		deck.Card{Value: 9, Suit: "spades"},
// 	}
// 	playerCards := make(map[*Client][]deck.Card)

// 	player1 := &Client{}
// 	player2 := &Client{}

// 	playerCards[player1] = []deck.Card{deck.Card{Value: 2, Suit: "spades"}, deck.Card{Value: 6, Suit: "clubs"}}
// 	playerCards[player2] = []deck.Card{deck.Card{Value: 2}, deck.Card{Value: 9}}
// 	fmt.Println("Test is running comparision function pair vs twoPairs")
// 	winners:= getWinnerHand(tableCards, playerCards)

// 	if winners[0]== player2 {
// 		fmt.Println("Test passed")
// 	} else {
// 		fmt.Println("Test failed higher pair should beat win")
// 	}
// }

// func testHandCompareTwoPairsToTwoPairsKicker() {
// 	tableCards := []deck.Card{
// 		deck.Card{Value: 2, Suit: "clubs"},
// 		deck.Card{Value: 4, Suit: "clubs"},
// 		deck.Card{Value: 2, Suit: "spades"},
// 		deck.Card{Value: 8, Suit: "hearts"},
// 		deck.Card{Value: 9, Suit: "spades"},
// 	}
// 	playerCards := make(map[*Client][]deck.Card)

// 	player1 := &Client{}
// 	player2 := &Client{}

// 	playerCards[player1] = []deck.Card{deck.Card{Value: 4, Suit: "spades"}, deck.Card{Value: 6, Suit: "clubs"}}
// 	playerCards[player2] = []deck.Card{deck.Card{Value: 4}, deck.Card{Value: 10}}
// 	fmt.Println("Test is running comparision two pairs kicker")
// 	winners:= getWinnerHand(tableCards, playerCards)

// 	if winners[0]== player1 {
// 		fmt.Println("Test passed")
// 	} else {
// 		fmt.Println("Test failed higher kicker should win")
// 	}
// }
