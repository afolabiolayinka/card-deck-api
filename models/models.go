package views

import "strconv"
import "strings"

import "api.card-deck/utils"


// `Card` object / model

type Card struct {
	Value string 	`json:"value"`
	Suit	string 	`json:"suit"`
	Code  string 	`json:"code"`
}

// `Deck` object / model

type Deck struct {
	Deck_id		string 	`json:"deck_id"`
	Shuffled	bool 		`json:"shuffled"`
	Remaining int 		`json:"remaining"`
	Cards 		[]Card 	`json:"cards,omitempty"`
}


// method to create card suits
// - `cards_query` (options) can be passed to create specific cards in deck
// - `cards_query` is retrieved through positional argument `args` on index 0 / position 0 

func CreateSuit (card_type string, args ...string) []Card {
	var __deck_cards = []string {"KING", "QUEEN", "JOKER"}
	var lettered_cards = []string {"KING", "QUEEN", "JOKER", "ACE"}
	var card_ordinary_length int = 9

	for i := card_ordinary_length; i >= 1; i-- {
		__deck_cards = append(__deck_cards, strconv.FormatInt(int64(i + 1), 10))
	}
	__deck_cards = append(__deck_cards, "ACE")

	var deck_cards = []Card {}

	for card := 0; card < len(__deck_cards); card++ {
		_value := strings.ToUpper(__deck_cards[card])
		_suit := strings.ToUpper(card_type)
		_code := _value + _suit[0:1]

		if (utils.StringInSlice(lettered_cards, __deck_cards[card])) {
			_code = _value[0:1] + _suit[0:1]
		}

		if (len(args) > 0) {
			cards_query := strings.Split(args[0], ",")

			if (utils.StringInSlice(cards_query, _code)) {
				deck_cards = append(
					deck_cards,
					Card {
						Value: _value,
						Suit: _suit,
						Code: _code,
					},
				)
			}
		} else {
			deck_cards = append(
				deck_cards,
				Card {
					Value: _value,
					Suit: _suit,
					Code: _code,
				},
			)
		}
	}

	return deck_cards
}
