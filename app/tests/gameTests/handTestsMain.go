package gameTests

import (
	"fmt"
	"riggedstars/app/deck"
	"riggedstars/app/game"
)

func getWinnerHand(tableCards []deck.Card, clientCards map[*game.Client][]deck.Card) []*game.Client {
	return game.GetWinnerHand(tableCards, clientCards)
}

func RunTests() {
	tests := []func() (string, bool, string){
		testHistogramHandCompareThreeOfAKindVsFullHouse,
		testHistogramHandCompareThreeOfAKindVsPair,
		testHistogramHandCompareFourOfAKindVsPair,
		testHistogramHandCompareFourOfAKindVsThreeOfAKind,
		testHistogramHandCompareFourOfAKindVsFullHouse,
		testHandComparePairVsStraight,
		testHandComparePairVsStraight2,
		testHandComparePairVsStraight3,
		testHandComparePairVsStraight4,
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
		testHandComparePairVsPair3,
		testHandComparePairVsPairTie,
		testHandCompareTwoPairVsTwoPair,
		testHandCompareTwoPairVsTwoPair2,
		testHandCompareTwoPairVsTwoPairTie,
		testHandCompareThreeOfAKindVsThreeOfAKind,
		testHandCompareThreeOfAKindVsThreeOfAKind2,
		testHandCompareThreeOfAKindVsThreeOfAKindTie,
		testHandCompareFourOfAKindVsFourOfAKind,
		testHandCompareFourOfAKindVsFourOfAKind2,
		testHandCompareFourOfAKindVsFourOfAKindTie,
		testHandCompareFullHouseVsFullHouse,
		testHandCompareFullHouseVsFullHouse2,
		testHandCompareFullHouseVsFullHouseTie,
		testHandCompareStraightVsStraight,
		testHandCompareStraightVsStraight2,
		testHandCompareStraightVsStraight3,
		testHandCompareStraightVsStraightTie,
		testHandCompareFlushVsFlush,
		testHandCompareFlushVsFlush2,
		testHandCompareFlushVsFlushTie,
		testHandCompareFlushVsFlushTie2,
		testHandCompareStraightFlushVsStraightFlush,
		testHandCompareStraightFlushVsStraightFlush2,
		testHandCompareStraightFlushVsStraightFlush3,
		testHandCompareStraightFlushVsStraightFlushTie,
		testHandCompareStraightFlushVsStraightFlushTie2,
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
