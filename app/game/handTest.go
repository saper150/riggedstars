package game

import (
	"fmt"
	"riggedstars/app/deck"
)

func RunTests() {
	testHistogramHandCompareThreeOfAKindVsFullHouse()
	testHistogramHandCompareThreeOfAKindVsPair()
	testHistogramHandCompareFourOfAKindVsPair()
	testHistogramHandCompareFourOfAKindVsThreeOfAKind()
	testHistogramHandCompareFourOfAKindVsFullHouse()
	testHandComparePairVsStraight()
	testHandComparePairVsFlush()
	testHandCompareStraightVsFlush()
	testHandCompareFlushVsStraightFlush()
	testHandCompareStraightVsStraightFlush()
	testHandCompareHighCardvsHighCardTier1()
	testHandCompareHighCardvsHighCardTier2()
	testHandCompareHighCardvsHighCardTier3()
	testHandCompareHighCardvsHighCardTier4()
	testHandCompareHighCardvsHighCardTier5()
	testHandCompareHighCardvsHighCardTie()
	testHandCompareHighCardvsHighCardAce1()
}

func testHistogramHandCompareThreeOfAKindVsPair() {
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
	fmt.Println("Test is running comparision function pair vs three of a kind")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}

	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed three of a kind should beat pair")
	}
}

func testHistogramHandCompareFourOfAKindVsPair() {
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
	fmt.Println("Test is running comparision function pair vs four of a kind")
	winners := getWinnerHand(tableCards, playerCards)

	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed four of a kind should beat pair")
	}
}

func testHistogramHandCompareFourOfAKindVsThreeOfAKind() {
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
	fmt.Println("Test is running comparision function three of a kind vs Four of a kind")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed Four of a kind should beat three of a kind")
	}
}

func testHistogramHandCompareFourOfAKindVsFullHouse() {
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
	fmt.Println("Test is running comparision function full house vs four of a kind")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed four of a kind should beat full house")
	}
}
func testHistogramHandCompareThreeOfAKindVsFullHouse() {
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
	fmt.Println("Test is running comparision function full house vs three of a kind")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed full house should beat three of a kind")
	}
}

func testHistogramHandComparePairVsHighCard() {
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
	fmt.Println("Test is running comparision function pair vs high card")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed pair should beat high card")
	}
}

func testHandComparePairVsStraight() {
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
	fmt.Println("Test is running comparision function pair vs straight")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed straight should beat pair")
	}
}

func testHandComparePairVsFlush() {
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
	fmt.Println("Test is running comparision function pair vs flush")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed flush should beat pair")
	}
}

func testHandCompareStraightVsFlush() {
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
	fmt.Println("Test is running comparision function straight vs flush")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed flush should beat straight")
	}
}

func testHandCompareFlushVsStraightFlush() {
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
	fmt.Println("Test is running comparision function flush vs straightflush")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed straightflush should beat flush")
	}
}

func testHandCompareStraightVsStraightFlush() {
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
	fmt.Println("Test is running comparision function pair vs straightflush")
	winners := getWinnerHand(tableCards, playerCards)

	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed straightflush should beat pair")
	}
}

func testHandCompareHighCardvsHighCardTier1() {
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
	fmt.Println("Test is running comparision function high card vs high card tier 1")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card Ace should be high card 5")
	}
}

func testHandCompareHighCardvsHighCardTier2() {
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
	fmt.Println("Test is running comparision function high card vs high card tier 2")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card King should be high card 5")
	}
}

func testHandCompareHighCardvsHighCardTier3() {
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
	fmt.Println("Test is running comparision function high card vs high card tier 3")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card Queen should be high card 5")
	}
}

func testHandCompareHighCardvsHighCardTier4() {
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
	fmt.Println("Test is running comparision function high card vs high card tier 4")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card Jack should be high card 5")
	}
}

func testHandCompareHighCardvsHighCardTier5() {
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
	fmt.Println("Test is running comparision function high card vs high card tier 5")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player1 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card 7 should be high card 5")
	}
}

func testHandCompareHighCardvsHighCardTie() {
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
	fmt.Println("Test is running comparision function high card vs high card tie")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) == 2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed expected a tie")
	}
}

func testHandCompareHighCardvsHighCardAce1() {
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
	fmt.Println("Test is running comparision function high card vs high card with Ace")
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		fmt.Println("Should be one winner")
	}
	if winners[0] == player2 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed high card Ace should be high card 6")
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
