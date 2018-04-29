package gameTests

import (
	"riggedstars/app/deck"
	"riggedstars/app/game"
)

func testHandCompareHighCardvsHighCardTier1() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 6, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "clubs"},
		deck.Card{Value: 11, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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

func testHandComparePairVsPair3() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 12}}
	name := "comparision function pair vs pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed pair of 1 should beat pair of 12"
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
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

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

func testHandCompareTwoPairVsTwoPair() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 12, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 12}, deck.Card{Value: 3}}
	name := "comparision function two pair vs two pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed two pairs 10 & 12 with high card Ace should beat two pairs 12 & 10 with high card 13"
	}
}

func testHandCompareTwoPairVsTwoPair2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 13, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 3}, deck.Card{Value: 7}}
	name := "comparision function two pair vs two pair"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed two pairs 13 & 10 with 12 high should beat two pairs 3 & 7 with 13 high "
	}
}

func testHandCompareTwoPairVsTwoPairTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 4, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function two pair vs two pair tie"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed two pairs 13 & 12 with 10 high should tie with two pairs 13 & 12 with 10 high"
	}
}

func testHandCompareThreeOfAKindVsThreeOfAKind() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 4, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 4}, deck.Card{Value: 2}}
	name := "comparision function Three Of A Kind Vs Three Of A Kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Three of a kind 10 Ace high should beat Three of a kind 10 Queen high"
	}
}

func testHandCompareThreeOfAKindVsThreeOfAKind2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 7, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 10, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 7}, deck.Card{Value: 7}}
	name := "comparision function Three Of A Kind Vs Three Of A Kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Three of a kind 10 should beat Three of a kind 7"
	}
}

func testHandCompareThreeOfAKindVsThreeOfAKindTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 4, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function Three Of A Kind Vs Three Of A Kind Tie"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Three of a kind 13 - 1,12 high should tie with Three of a kind 13 - 1,12 high"
	}
}

func testHandCompareFourOfAKindVsFourOfAKind() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 4, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 4}, deck.Card{Value: 2}}
	name := "comparision function Four Of A Kind Vs Four Of A Kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Four of a kind 10 Ace high should beat Four of a kind 10 Queen high"
	}
}

func testHandCompareFourOfAKindVsFourOfAKind2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 10, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 3}, deck.Card{Value: 3}}
	name := "comparision function Four Of A Kind Vs Four Of A Kind"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Four of a kind 10 should beat Four of a kind 3"
	}
}

func testHandCompareFourOfAKindVsFourOfAKindTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 4, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function Four Of A Kind Vs Four Of A Kind Tie"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Four of a kind 13 - 1 high should tie with Four of a kind 13 - 1 high"
	}
}

func testHandCompareFullHouseVsFullHouse() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 1, Suit: "hearts"},
		deck.Card{Value: 12, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 4, Suit: "clubs"}, deck.Card{Value: 1, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 4}, deck.Card{Value: 12}}
	name := "comparision function Full House vs Full House"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Full House 10 of 1 should beat Full House 10 of 12 "
	}
}

func testHandCompareFullHouseVsFullHouse2() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 10, Suit: "clubs"},
		deck.Card{Value: 10, Suit: "hearts"},
		deck.Card{Value: 3, Suit: "clubs"},
		deck.Card{Value: 12, Suit: "hearts"},
		deck.Card{Value: 2, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 12, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 3}, deck.Card{Value: 3}}
	name := "comparision function Full House vs Full House"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, false, "Failed - Should be one winner"
	}
	if winners[0] == player1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Full House 10 of 12 should beat Full House 3 of 10"
	}
}

func testHandCompareFullHouseVsFullHouseTie() (string, bool, string) {
	tableCards := []deck.Card{
		deck.Card{Value: 1, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 13, Suit: "clubs"},
		deck.Card{Value: 13, Suit: "hearts"},
		deck.Card{Value: 1, Suit: "spades"},
	}
	playerCards := make(map[*game.Client][]deck.Card)

	player1 := &game.Client{}
	player2 := &game.Client{}

	playerCards[player1] = []deck.Card{deck.Card{Value: 10, Suit: "clubs"}, deck.Card{Value: 4, Suit: "clubs"}}
	playerCards[player2] = []deck.Card{deck.Card{Value: 10}, deck.Card{Value: 2}}
	name := "comparision function Full House vs Full House"
	winners := getWinnerHand(tableCards, playerCards)
	if len(winners) != 1 {
		return name, true, "Test passed"
	} else {
		return name, false, "Test failed Full House 13 of 1 should tie with Full House 13 of 1"
	}
}
